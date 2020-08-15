package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Pusher-Labs/midgard-graphql/graph"
	"github.com/Pusher-Labs/midgard-graphql/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r, err := graph.NewResolver()
	if err != nil {
		fmt.Errorf("error creating resolver: %v", err)
	}
	
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: r}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
