package main

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/jihadable/sticker-be/graph"
	"github.com/jihadable/sticker-be/utils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.local")
	if err != nil {
		panic("Failed to read .env file: " + err.Error())
	}

	app := fiber.New()

	app.All("/graphql", func(c *fiber.Ctx) error {
		handler := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
		handler.AddTransport(transport.POST{})
		handler.AddTransport(transport.GET{})

		authHeader := c.Get("Authorization")
		ctx := context.WithValue(c.Context(), utils.AuthHeader{}, authHeader)

		return adaptor.HTTPHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(ctx)
			handler.ServeHTTP(w, r)
		}))(c)
	})

	app.Listen(":3000")
}
