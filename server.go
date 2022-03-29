package main

import (
	"elections-api/database"
	"elections-api/graph"
	"elections-api/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}

	database.ConnectDB()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
