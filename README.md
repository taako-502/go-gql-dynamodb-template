# Go GraphQL and DynamoDB and Serverless Template

## Local Development

```bash
go run handler/main.go -migrate
docker-compose up -d
```

## GraphQL Generation

```bash
cd handler
go run github.com/99designs/gqlgen generate
```

## Deploy

```bash
make && serverless deploy
```
