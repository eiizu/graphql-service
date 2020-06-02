package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gitlab.srconnect.io/acuevas/graphql-server/graph/dataloader"
	"gitlab.srconnect.io/acuevas/graphql-server/graph/generated"
	"gitlab.srconnect.io/acuevas/graphql-server/graph/resolver"
	"gitlab.srconnect.io/acuevas/graphql-server/graph/store"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := store.New()
	dl := dataloader.New(db)

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: resolver.New(db),
			},
		))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", dl.Middleware(srv))

	log.Printf("server running on localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
