package main

import (
	"errors"
	"graphql/graph"
	"graphql/infrastructure/datastore"
	"graphql/infrastructure/graphql"
	"graphql/infrastructure/router"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	client, err := datastore.NewRedisClient("20.5.40.162:8000")
	if !errors.Is(err, nil) {
		log.Fatalln(err)
	}
	defer client.Close()

	r := graph.NewResolver(client)
	r.SubscribeRedis()
	srv := graphql.NewGraphQLServer(r)

	e := router.NewRouter(echo.New(), srv)
	e.Logger.Fatal(e.Start(":8080"))
}
