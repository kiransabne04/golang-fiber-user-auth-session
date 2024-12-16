package middleware

import (
	"errors"
	"fiber-user-auth-session/internal/session"
	"fiber-user-auth-session/pkg"
	"time"

	"github.com/gofiber/fiber/v2"
)

const maxIdleTimeout = 15 * time.Minute

func SessionPGMiddleware (repo *session.SessionRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionID := c.Cookies("session_id")
		if sessionID == "" {
			return pkg.ErrorJSON(c, errors.New("invalid request payload"), fiber.StatusUnauthorized)
		}

		// Validate session
		session, err := repo.GetSessionByID(c.Context(), sessionID)
		if err != nil || !session.IsActive {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "Invalid or expired session",
			})
		}

		// Check IP address
		currentIP := c.IP()
		if session.IPAddress != currentIP {
			repo.InvalidateSession(c.Context(), sessionID) // Invalidate session on IP mismatch
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "IP address mismatch",
			})
		}

		// Check last activity
		if time.Since(session.LastActivity) > maxIdleTimeout {
			repo.InvalidateSession(c.Context(), sessionID) // Invalidate session on idle timeout
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "Session has expired due to inactivity",
			})
		}

		// Update last activity
		err = repo.UpdateLastActivity(c.Context(), sessionID, time.Now())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Failed to update session activity",
			})
		}
		
		return c.Next()
	}
}