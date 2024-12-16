package user

import (
	"context"
	"errors"
	"fiber-user-auth-session/pkg"

	"github.com/gofiber/fiber/v2"
)

// RegisterUser handles user registration requests.
func RegisterUser(c *fiber.Ctx, services *UserService) error {
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
	id, err := services.RegisterUser(context.Background(), req.Name, req.Email, req.Password)
	if err != nil {
		return pkg.ErrorJSON(c, err, fiber.StatusInternalServerError)
	}

	return pkg.SuccessJSON(c, "User registered successfully", fiber.Map{
		"user_id": id,
	})
}

func TestUser(c *fiber.Ctx, services *UserService) error {
	
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": "kks"}})
}

func LoginHandler(c *fiber.Ctx, services *UserService) error {
	
	var signInReq SignInInput

	if err := c.BodyParser(&signInReq); err != nil {
		return pkg.ErrorJSON(c, errors.New("invalid request payload"))
	}

	// Call the service to authenticate the user
	user, err := services.LoginUser(c.Context(), signInReq.Email, signInReq.Password)
	if err != nil {
		return pkg.ErrorJSON(c, err, fiber.StatusUnauthorized)
	}

	// set a session cookie (for web) or return JWT tokens for mobile/frameworks
	isWebClient := c.Get("User-Agent") != "" // Example check; refine based on use case

	if isWebClient {
		//set session cookie for web
		sess, err := services.
	}
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
