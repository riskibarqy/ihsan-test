package http

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App, userHandler *UserHandler) {
	api := app.Group("/api")
	api.Get("/users/:id", userHandler.GetUser)
	api.Post("/users", userHandler.CreateUser)
}
