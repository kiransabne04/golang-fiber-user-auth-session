package pkg

import "github.com/gofiber/fiber/v2"

type JSONResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// WriteJSON sends a JSON response with the given status code and headers.
func WriteJSON(c *fiber.Ctx, status int, data JSONResponse) error {
	c.Status(status)
	return c.JSON(data)
}

// ErrorJSON sends an error response with the given status code.
func ErrorJSON(c *fiber.Ctx, err error, status ...int) error {
	statusCode := fiber.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	response := JSONResponse{
		Error:   true,
		Message: err.Error(),
	}
	return WriteJSON(c, statusCode, response)
}

// SuccessJSON sends a success response with optional data.
func SuccessJSON(c *fiber.Ctx, message string, data interface{}) error {
	response := JSONResponse{
		Error:   false,
		Message: message,
		Data:    data,
	}
	return WriteJSON(c, fiber.StatusOK, response)
}