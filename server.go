package main

import (
	"log"
	"net/http"
	"os"

	"gograb/graph"
	"gograb/graph/generated"
	"gograb/internal/auth"
	_ "gograb/internal/auth"
	database "gograb/internal/pkg/db/mysql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(auth.Middleware())

	database.InitDB()
	defer database.CloseDB()
	database.Migrate()
	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/", playground.Handler("GraphQL", "/query"))
	router.Handle("/query", server)
	// base route url -> move to grabhQL interface
	log.Printf("connect to http://localhost:%s/ for GraphQL", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
