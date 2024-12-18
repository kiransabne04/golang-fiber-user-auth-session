package routers

import (
	"fiber-user-auth-session/internal/auth"

	"github.com/gofiber/fiber/v2"
)

// func authRoutes(v1 fiber.Router, services *internal.AppServices) {
// 	authGroup := v1.Group("/auth")
// 	authGroup.Get("/test", func(c *fiber.Ctx) error {
// 		return authGroup.LoginHandler(c, services)
// 	})
// }

func RegisterAuthRoutes(v1 fiber.Router, authHandler *auth.AuthHandler) {
	authGroup := v1.Group("/auth")
	authGroup.Post("/login", authHandler.LoginHandler)
}
