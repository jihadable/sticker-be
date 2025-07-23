package handler

import (
	"context"
	"fmt"
	"net/http"

	graphqlHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jihadable/sticker-be/config"
	"github.com/jihadable/sticker-be/graph"
	"github.com/jihadable/sticker-be/services"
	"github.com/jihadable/sticker-be/validators"
	"github.com/joho/godotenv"
)

// Handler is the main entry point of the application. Think of it like the main() method
func Handler(w http.ResponseWriter, r *http.Request) {
	// This is needed to set the proper request path in `*fiber.Ctx`
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

// building the fiber application
func handler() http.HandlerFunc {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Warning: .env file not found")
	}

	app := fiber.New()
	app.Use(cors.New(cors.ConfigDefault))

	db := config.DB()
	redis := config.Redis()
	pusher := config.NewPusher()

	handler := graphqlHandler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		UserService:            services.NewUserService(db, redis),
		ProductService:         services.NewProductService(db, redis),
		CustomProductService:   services.NewCustomProductService(db, redis),
		CategoryService:        services.NewCategoryService(db, redis),
		ProductCategoryService: services.NewProductCategoryService(db, redis),
		CartService:            services.NewCartService(db, redis),
		CartProductService:     services.NewCartProductService(db, redis),
		OrderService:           services.NewOrderService(db, redis, services.NewOrderProductService(db, redis)),
		ConversationService:    services.NewConversationService(db, redis),
		MessageService:         services.NewMessageService(db, redis),
		NotificationService:    services.NewNotificationService(db, redis),
		Pusher:                 pusher,
	}}))

	handler.Use(extension.Introspection{})
	handler.AddTransport(transport.POST{})
	handler.AddTransport(transport.GET{})
	handler.AddTransport(transport.MultipartForm{})
	handler.AddTransport(transport.Options{})

	app.All("/graphql", func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		ctx := context.WithValue(c.Context(), validators.AuthHeader, authHeader)

		return adaptor.HTTPHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(ctx)
			handler.ServeHTTP(w, r)
		}))(c)
	})

	app.Get("/docs", adaptor.HTTPHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../docs/index.html")
	})))

	return adaptor.FiberApp(app)
}
