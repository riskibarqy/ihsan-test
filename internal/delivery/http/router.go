package http

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App, userHandler *UserHandler) {
	api := app.Group("/api")

	// users
	api.Get("/users/:id", userHandler.GetUser)          // get user by id
	api.Post("/users/daftar", userHandler.RegisterUser) // regiser user
}
