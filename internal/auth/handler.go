package auth

import (
	"errors"
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

func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	type request struct {
		RefreshToken string `json:"refresh_token"`
	}

	var req request
	if err := c.BodyParser(&req); err != nil {
		return pkg.ErrorJSON(c, err, fiber.StatusBadRequest)
	}

	// parse & Validate refresh token
	claims, err := pkg.ParseToken(req.RefreshToken)
	if err != nil || claims == nil {
		return pkg.ErrorJSON(c, errors.New("invalid or expired refresh token"), fiber.StatusUnauthorized)
	}

	// check if the session is still alive
	sessionID := claims.SessionID
	session, err := h.AuthService.SessionRepo.GetSessionByID(c.Context(), sessionID)
	if err != nil || !session.IsActive {
		return pkg.ErrorJSON(c, errors.New("session is inactive or does not exists"), fiber.StatusUnauthorized)
	}

	//generate a new access token
	accessToken, err := pkg.GenerateAccessToken(sessionID, session.PersonID, h.AuthService.SecretKey)
	if err != nil {
		return pkg.ErrorJSON(c, errors.New("failed to generate access token"), fiber.StatusInternalServerError)
	}

	return pkg.SuccessJSON(c, "Token refreshed successfully", fiber.Map{
		"access_token": accessToken,
	})
}

func (h *AuthHandler) LogoutHandler(c *fiber.Ctx) error {
	sessionID := c.Cookies("session_id")
	if sessionID == "" {
		return pkg.ErrorJSON(c, errors.New("missing session id"), fiber.StatusUnauthorized)
	}

	err := h.AuthService.SessionRepo.InvalidateSession(c.Context(), sessionID)
	if err != nil {
		return pkg.ErrorJSON(c, errors.New("failed to invalidate session"), fiber.StatusInternalServerError)
	}

	c.ClearCookie("session_id")
	return pkg.SuccessJSON(c, "Logged out successfully", nil)
}
