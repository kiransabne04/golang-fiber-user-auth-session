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

		var sessionID string
		var isTokenBased bool

		// Check Authorization header for mobile/JS clients
		token := c.Get("Authorization")
		if token != "" {
			isTokenBased = true
			claims, err := pkg.ParseToken(token)
			if err != nil {
				return pkg.ErrorJSON(c, errors.New("invalid or expired token"), fiber.StatusUnauthorized)
			}
			sessionID = claims.SessionID
		} else {
			// Check session_id cookie for web clients
			sessionID = c.Cookies("session_id")
			if sessionID == "" {
				return pkg.ErrorJSON(c, errors.New("missing session_id or token"), fiber.StatusUnauthorized)
			}
		}

		// Retrieve session from the database
		session, err := sessionService.GetSessionByID(c.Context(), sessionID)
		if err != nil || !session.IsActive {
			return pkg.ErrorJSON(c, errors.New("inactive or invalid session"), fiber.StatusUnauthorized)
		}

		// Validate IP address for web clients
		if !isTokenBased {
			currentIP := c.IP()
			if session.IPAddress != "" && session.IPAddress != currentIP {
				_ = sessionService.InvalidateSession(c.Context(), sessionID) // Invalidate session on IP mismatch
				return pkg.ErrorJSON(c, errors.New("IP address mismatch"), fiber.StatusUnauthorized)
			}
		}

		// Check for idle timeout
		if time.Since(session.LastActivity) > idleTimeout {
			_ = sessionService.InvalidateSession(c.Context(), sessionID) // Invalidate session on idle timeout
			return pkg.ErrorJSON(c, errors.New("session expired due to inactivity"), fiber.StatusUnauthorized)
		}

		// Update last activity
		err = sessionService.UpdateLastActivity(c.Context(), sessionID)
		if err != nil {
			return pkg.ErrorJSON(c, errors.New("failed to update session activity"), fiber.StatusInternalServerError)
		}

		// Set user ID in context for downstream handlers
		c.Locals("user_id", session.PersonID)
		return c.Next()

	}
}
