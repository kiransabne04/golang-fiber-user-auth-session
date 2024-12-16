package routers

import (
	"fiber-user-auth-session/internal"
	"fiber-user-auth-session/internal/user"

	"github.com/gofiber/fiber/v2"
)

func userRoutes(v1 fiber.Router, services *internal.AppServices) {
	userGroup := v1.Group("/user")

	// Handlers for user routes
	//userGroup.Get("/test", handlers.UserTest(services))
	//userGroup.Post("/register", handlers.UserRegister(services))
	userGroup.Get("/test", func(c *fiber.Ctx) error {
		return user.TestUser(c, services.UserService)
	})
}