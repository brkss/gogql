package main

import (
	//"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/brkss/gogql/directive"
	"github.com/brkss/gogql/graph"
	"github.com/brkss/gogql/utils"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config : ", err);
	}

	//con, err := sql.Open(config.DBDriver, config.DBSource)
	fmt.Printf("config : %+v\n", config)

	c := graph.Config{Resolvers: &graph.Resolver{}}
	c.Directives.Binding = directive.Binding

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
