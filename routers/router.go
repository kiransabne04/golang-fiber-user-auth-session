package routers

import (
	"fiber-user-auth-session/internal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func SetupRouters(app *fiber.App, services *internal.AppServices) {
	app.Use(
		requestid.New(),
		cors.New(),
		recover.New(),
	)

	//api versioning
	v1 := app.Group("/v1")

	// v1.Route("/user", func(router fiber.Router) {
	// 	router.Get("/test", handlers.Test)
	// })
	//

	// user routes
	userRoutes(v1, services)
}