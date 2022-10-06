package routes

import (
	"app/handlers"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {

	r := e.Group("/")

	e.GET("/playground", func(c echo.Context) error {
		playground.Handler("GraphQL", "/query").ServeHTTP(c.Response(), c.Request())
		return nil
	})

	r.Any("/", handlers.PlaygroundHandler())
	r.Any("ws", handlers.PlaygroundWsHandler())

	r.Any("query", handlers.GraphqlHandler())
	r.Any("subscription", handlers.GraphqlWsHandler())
}
