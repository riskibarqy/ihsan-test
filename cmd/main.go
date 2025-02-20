package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/riskibarqy/ihsan-test/internal/delivery/http"
	"github.com/riskibarqy/ihsan-test/internal/repository"
	"github.com/riskibarqy/ihsan-test/internal/usecase"
	"github.com/riskibarqy/ihsan-test/pkg/database"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Connect to database
	db, err := database.ConnectPostgres()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Repository, Usecase, and Handlers
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := http.NewUserHandler(userUsecase)

	// Setup Routes
	http.SetupRoutes(app, userHandler)

	// Start server
	log.Fatal(app.Listen(":8080"))
}
