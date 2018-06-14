package main

import (
	"log"
	"net/http"

	"github.com/skimke/graphql-app/views"
)

func main() {
	http.HandleFunc("/graphiql", views.ServeGraphiQL)

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
