package auth

import "github.com/gofiber/fiber/v2"

type AuthHandler struct {
	AuthService *AuthService
}

func NewAuthHandler(authService *AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {

}
