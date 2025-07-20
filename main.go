package main

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

func main() {
	err := godotenv.Load(".env.local")
	if err != nil {
		panic("Failed to read .env file: " + err.Error())
	}

	app := fiber.New()
	app.Use(cors.New(cors.ConfigDefault))

	db := config.DB()
	redis := config.Redis()

	app.All("/graphql", func(c *fiber.Ctx) error {
		handler := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
			UserService:                  services.NewUserService(db, redis),
			ProductService:               services.NewProductService(db, redis),
			CustomProductService:         services.NewCustomProductService(db, redis),
			CategoryService:              services.NewCategoryService(db, redis),
			ProductCategoryService:       services.NewProductCategoryService(db, redis),
			CartService:                  services.NewCartService(db, redis),
			CartProductService:           services.NewCartProductService(db, redis),
			OrderService:                 services.NewOrderService(db, redis),
			OrderProductService:          services.NewOrderProductService(db, redis),
			ConversationService:          services.NewConversationService(db, redis),
			MessageService:               services.NewMessageService(db, redis),
			NotificationService:          services.NewNotificationService(db, redis),
			NotificationRecipientService: services.NewNotificationRecipientService(db, redis),
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

	// app.Get("/", adaptor.HTTPHandler(playground.Handler("GraphQL Playground", "/graphql")))

	app.Listen(":3000")
}

// PostUser is the resolver for the post_user field.
// func (r *mutationResolver) PostUser(ctx context.Context, name string, email string, password string, phone string, address string) (*model.Auth, error) {
// 	user, err := r.UserService.AddUser(&models.User{
// 		Name:     name,
// 		Email:    email,
// 		Password: password,
// 		Phone:    &phone,
// 		Address:  &address,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	_, err = r.CartService.AddCart(&models.Cart{CustomerId: user.Id})
// 	if err != nil {
// 		return nil, err
// 	}

// 	admin, err := r.UserService.GetAdmin()
// 	if err != nil {
// 		return nil, err
// 	}

// 	conversation, err := r.ConversationService.AddConversation(&models.Conversation{
// 		CustomerId: user.Id,
// 		AdminId:    admin.Id,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	newMessage, err := r.MessageService.AddMessage(&models.Message{
// 		ConversationId: conversation.Id,
// 		SenderId:       admin.Id,
// 		Message:        "Selamat datang di stikerin",
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	notification, err := r.NotificationService.AddNotification(&models.Notification{
// 		Title:   "Selamat datang di stikerin",
// 		Message: newMessage.Message,
// 		Type:    "new_message",
// 	}, user.Id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = config.MessageTrigger("new_message", newMessage)
// 	if err != nil {
// 		return nil, err
// 	}

// 	pushNotification := models.PushNotification{
// 		UserId:  user.Id,
// 		Type:    notification.Type,
// 		Title:   notification.Title,
// 		Message: notification.Message,
// 		IsRead:  false,
// 	}
// 	err = config.NotificationTrigger("new_message_notification_"+user.Id, pushNotification)
// 	if err != nil {
// 		return nil, err
// 	}

// 	token, err := utils.GenerateJWT(user.Id, user.Role)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &model.Auth{User: mapper.DBUserToGraphQLUser(user), Token: *token}, nil
// }
