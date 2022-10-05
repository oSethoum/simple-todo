package routes

import (
	"app/handlers"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {

	r := e.Group("/")

	r.Any("", handlers.PlaygroundHandler())
	r.Any("ws", handlers.PlaygroundWsHandler())

	r.Any("query", handlers.GraphqlHandler())
	r.Any("subscriptions", handlers.GraphqlWsHandler())
}
