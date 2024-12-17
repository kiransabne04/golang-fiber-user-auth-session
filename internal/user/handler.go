package user

import (
	"context"
	"errors"
	"fiber-user-auth-session/pkg"
	"log"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService *UserService
}

func NewUserHandler(services *UserService) *UserHandler {
	return &UserHandler{UserService: services}
}

// RegisterUser handles user registration requests.
func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	type request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req request
	if err := c.BodyParser(&req); err != nil {
		return pkg.ErrorJSON(c, errors.New("invalid request payload"))
	}

	// Call the service to register the user
	id, err := h.UserService.RegisterUser(context.Background(), req.Name, req.Email, req.Password)
	if err != nil {
		return pkg.ErrorJSON(c, err, fiber.StatusInternalServerError)
	}

	return pkg.SuccessJSON(c, "User registered successfully", fiber.Map{
		"user_id": id,
	})
}

func (h *UserHandler) TestUser(c *fiber.Ctx) error {
	log.Println("testUser -> ")
	//return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": "kks"}})
	return pkg.SuccessJSON(c, "asgeraraewr", fiber.Map{"user": "kks"})
}

// varsha shashank madgunik - ambernath - 9869486121
// ankita manmath mahajan - 9049229828
// shivani suresh kahlekar - 9011470111
// archana chokanpale - 9833698526
// ashvini shivajiappa nagathne - 9673905225
// dakshyani wadgaokar - 9822901193

// shraddha malipatil - 8390815542
// vaishnavi pradip nikhle - 9702577988
// Neha dhulshete - 7385861678, 7887554543, 8459455787
// Namrata Khake - 9421085651, 9518772187
