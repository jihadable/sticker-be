package graphql

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jihadable/sticker-be/config"
	"github.com/jihadable/sticker-be/graph"
	"github.com/jihadable/sticker-be/services"
	"github.com/jihadable/sticker-be/validators"
	"github.com/joho/godotenv"
)

func TestApp() *fiber.App {
	err := godotenv.Load("../../.env.local")
	if err != nil {
		panic("Failed to read .env file: " + err.Error())
	}

	app := fiber.New()
	app.Use(cors.New(cors.ConfigDefault))

	db := config.DB()
	redis := config.Redis()

	app.All("/graphql", func(c *fiber.Ctx) error {
		handler := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
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
		}}))

		handler.AddTransport(transport.POST{})
		handler.AddTransport(transport.GET{})
		handler.AddTransport(transport.MultipartForm{})

		authHeader := c.Get("Authorization")
		ctx := context.WithValue(c.Context(), validators.AuthHeader, authHeader)

		return adaptor.HTTPHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(ctx)
			handler.ServeHTTP(w, r)
		}))(c)
	})

	return app
}

var App = TestApp()
var CustomerJWT, AdminJWT string
var ProductId, CustomProductId string
var CategoryId string
var ConversationId string
var OrderId string
var CartId, CartProductId string
var NotificationId string
