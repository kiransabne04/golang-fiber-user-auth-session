package routers

import (
	"fiber-user-auth-session/internal/middleware"
	"fiber-user-auth-session/internal/session"
	"fiber-user-auth-session/internal/user"
	"time"

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
func RegisterUserRoutes(v1 fiber.Router, userHandler *user.UserHandler, sessionService *session.SessionService, secretKey []byte) {
	userGroup := v1.Group("/user")

	// User routes
	userGroup.Get("/test", userHandler.TestUser)
	userGroup.Post("/register", userHandler.RegisterUser)

	userGroup.Use(middleware.SessionValidationMiddleware(sessionService, secretKey, 15*time.Minute))

	userGroup.Get("/profile", userHandler.TestUser)

}
