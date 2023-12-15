package main

import (
	"context"
	"flag"
	"fmt"

	"go-gql-dynamodb-template/handler/graph"
	ddbmanager "go-gql-dynamodb-template/handler/graph/dynamodb"
	"go-gql-dynamodb-template/handler/internal/config"
	"go-gql-dynamodb-template/handler/internal/localserver"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"

	"github.com/guregu/dynamo"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var echoLambda *echoadapter.EchoLambda

// Defining the Graphql handler
func graphqlHandler(db *dynamo.DB) echo.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))

	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

// Handler is the main function called by AWS Lambda.
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if echoLambda == nil {
		log.Printf("Echo cold start")
		e := echo.New()
		hosts := []string{
			os.Getenv("FRONTEND_HOST_1"),
			os.Getenv("FRONTEND_HOST_2"),
			os.Getenv("FRONTEND_HOST_3"),
		}
		e.Use(config.SettingCorsForEcho(hosts))

		db := ddbmanager.New("")
		e.POST("/query", graphqlHandler(db))

		echoLambda = echoadapter.New(e)
	}

	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	// 環境変数読み込み
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	// コマンドライン引数をパース
	migrate := flag.Bool("migrate", false, "Run database migrations")
	flag.Parse()

	// ローカル環境で打鍵するときに使う
	// go run handler/main.go -migrate
	if *migrate {
		// DynamoDBの初期化
		endpoint := os.Getenv("MIGRATION_ENDPOINT")
		db := ddbmanager.New(endpoint)
		manager := localserver.DDBMnager{DB: db}

		// マイグレーション実行
		fmt.Println("Running migrations...")
		if err := manager.Migration(); err != nil {
			log.Fatalf("マイグレーションに失敗しました: %v", err)
			os.Exit(1)
		}
		fmt.Println("マイグレーションが完了しました。")
		return
	}

	if os.Getenv("ENVIRONMENT") == "local" {
		// ローカル環境
		endpoint := os.Getenv("DYNAMO_ENDPOINT")
		db := ddbmanager.New(endpoint)
		localserver.StartLocalServer(db)
	} else {
		// AWS Lambda
		lambda.Start(Handler)
	}
}
