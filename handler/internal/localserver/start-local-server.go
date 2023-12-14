package localserver

import (
	"go-gql-dynamodb-template/handler/graph"
	"go-gql-dynamodb-template/handler/internal/config"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/guregu/dynamo"
)

func StartLocalServer(db *dynamo.DB) {
	// Routerの設定
	router := chi.NewRouter()
	router.Use(config.SettingCrosForLocalServer().Handler)
	graphqlServer := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", graphqlServer)

	graphqlServerPort := os.Getenv("GRAPHQL_SERVER_PORT")
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", graphqlServerPort)
	if err := http.ListenAndServe(":"+graphqlServerPort, router); err != nil {
		panic(err)
	}
}
