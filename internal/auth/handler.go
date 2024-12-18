package auth

import (
	"fiber-user-auth-session/pkg"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthService *AuthService
}

func NewAuthHandler(authService *AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (h *AuthHandler) LoginHandler(c *fiber.Ctx) error {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req request
	if err := c.BodyParser(&req); err != nil {
		return pkg.ErrorJSON(c, err, fiber.StatusBadRequest)
	}

	accessToken, refreshToken, err := h.AuthService.LoginService(c.Context(), req.Email, req.Password)
	if err != nil {
		return pkg.ErrorJSON(c, err, fiber.StatusUnauthorized)
	}

	return pkg.SuccessJSON(c, "Login successful", fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
