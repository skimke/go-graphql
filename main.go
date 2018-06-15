package main

import (
	"log"
	"net/http"

	"github.com/skimke/graphql-app/graphapi/gql"
	"github.com/skimke/graphql-app/graphapi/resolvers"
	"github.com/vektah/gqlgen/handler"
)

func main() {
	schema := gql.MakeExecutableSchema(resolvers.New())

	http.Handle("/query", handler.GraphQL(schema))
	http.Handle("/graphiql", handler.Playground("Emojis", "/query"))

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
