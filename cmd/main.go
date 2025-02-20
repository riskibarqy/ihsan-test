package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-migrate/migrate/v4"

	"github.com/riskibarqy/ihsan-test/internal/config"
	"github.com/riskibarqy/ihsan-test/internal/delivery/http"
	"github.com/riskibarqy/ihsan-test/internal/repository"
	"github.com/riskibarqy/ihsan-test/internal/usecase"
	"github.com/riskibarqy/ihsan-test/pkg/database"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(healthcheck.New())

	cfg := config.LoadConfig()

	db, err := database.ConnectPostgres(cfg)
	if err != nil {
		log.Fatal(err, cfg.GetDatabaseDSN())
	}

	// Run Migrations before starting the app
	runMigrations(cfg)

	// Initialize Repository, Usecase, and Handlers

	// user
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := http.NewUserHandler(userUsecase)

	// user balance history
	userBalanceHistoryRepo := repository.NewUserBalanceHistoryRepository(db)
	userBalanceHistoryUsecase := usecase.NewUserBalanceHistoryUsecase(userBalanceHistoryRepo, userRepo)
	userBalanceHistoryHandler := http.NewUserBalanceHistoryHandler(userBalanceHistoryUsecase)

	// Setup Routes
	http.SetupRoutes(app, userHandler, userBalanceHistoryHandler)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", cfg.AppPort)))
}

// runMigrations running file migrations
func runMigrations(cfg *config.Config) {
	m, err := migrate.New(
		"file://migrations",
		cfg.GetDatabaseDSN(),
	)
	if err != nil {
		log.Fatal("Migration initialization failed: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migration failed: %v", err)
	}

	log.Info("Migrations applied successfully")
}
