package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	db "github.com/brkss/gogql/db/sqlc"
	"github.com/brkss/gogql/directive"
	"github.com/brkss/gogql/graph"
	"github.com/brkss/gogql/utils"
	_ "github.com/lib/pq"
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

	con, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database : ", err)
	}
	store := db.NewStore(con)
	

	c := graph.Config{Resolvers: &graph.Resolver{
		Config: config,
		Store: store,
	}}

	c.Directives.Binding = directive.Binding

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
