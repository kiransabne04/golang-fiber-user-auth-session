package routers

import (
	"fiber-user-auth-session/internal/user"

	"github.com/gofiber/fiber/v2"
)

// func userRoutes(v1 fiber.Router, userHandler *user.UserHandler) {
// 	userGroup := v1.Group("/user")

//		// Handlers for user routes
//		//userGroup.Get("/test", handlers.UserTest(services))
//		//userGroup.Post("/register", handlers.UserRegister(services))
//		userGroup.Get("/test", userHandler.TestUser)
//		userGroup.Post("/register", userHandler.RegisterUser)
//	}
func RegisterUserRoutes(v1 fiber.Router, userHandler *user.UserHandler) {
	userGroup := v1.Group("/user")

	// User routes
	userGroup.Get("/test", userHandler.TestUser)
	userGroup.Post("/register", userHandler.RegisterUser)
}
