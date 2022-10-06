package handlers

import (
	"app/graph/resolvers"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func PlaygroundHandler(c echo.Context) error {
	h := playground.Handler("GraphQL", "/query")
	h.ServeHTTP(c.Response(), c.Request())
	return nil
}

func PlaygroundWsHandler(c echo.Context) error {
	h := playground.Handler("GraphQL WS", "/subscription")
	h.ServeHTTP(c.Response(), c.Request())
	return nil
}

func GraphqlWsHandler(c echo.Context) error {
	h := handler.New(resolvers.ExecutableSchema())
	h.Use(extension.Introspection{})
	h.AddTransport(transport.POST{})
	h.AddTransport(&transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	h.ServeHTTP(c.Response(), c.Request())
	return nil
}

func GraphqlHandler(c echo.Context) error {
	h := handler.NewDefaultServer(resolvers.ExecutableSchema())
	h.Use(extension.Introspection{})
	h.ServeHTTP(c.Response(), c.Request())
	return nil
}
