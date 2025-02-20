package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/riskibarqy/ihsan-test/internal/config"
	"github.com/riskibarqy/ihsan-test/internal/delivery/http"
	"github.com/riskibarqy/ihsan-test/internal/repository"
	"github.com/riskibarqy/ihsan-test/internal/usecase"
	"github.com/riskibarqy/ihsan-test/pkg/database"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(healthcheck.New())

	cfg := config.LoadConfig()

	db, err := database.ConnectPostgres(cfg)
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
	log.Fatal(app.Listen(fmt.Sprintf(":%s", cfg.AppPort)))
}
