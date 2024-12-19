package middleware

import (
	"errors"
	"fiber-user-auth-session/internal/session"
	"fiber-user-auth-session/pkg"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SessionValidationMiddleware(sessionService *session.SessionService, secretKey []byte, idleTimeout time.Duration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return pkg.ErrorJSON(c, errors.New("missing authorization token"), fiber.StatusUnauthorized)
		}

		claims, err := pkg.ParseToken(token)
		if err != nil {
			return pkg.ErrorJSON(c, errors.New("invalid or expired token"), fiber.StatusUnauthorized)
		}

		// retrieve the session from database
		session, err := sessionService.GetSessionByID(c.Context(), claims.SessionID)
		if err != nil || !session.IsActive {
			return pkg.ErrorJSON(c, errors.New("inactive or invalid session"), fiber.StatusUnauthorized)
		}

		// check for idle timeout
		if time.Since(session.LastActivity) > idleTimeout {
			_ = sessionService.InvalidateSession(c.Context(), claims.SessionID)
			return pkg.ErrorJSON(c, errors.New("session timed out due to inactivity"), fiber.StatusUnauthorized)
		}

		//update last activity
		err = sessionService.UpdateLastActivity(c.Context(), claims.SessionID)
		if err != nil {
			return pkg.ErrorJSON(c, errors.New("failed to update session activity"), fiber.StatusInternalServerError)
		}

		// Set user ID in context for downstream handlers
		c.Locals("user_id", session.PersonID)
		return c.Next()
	}
}
