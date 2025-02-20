package http

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App, userHandler *UserHandler, userBalanceHistoryHandler *UserBalanceHistoryHandler) {
	api := app.Group("/api")

	// users
	api.Get("/users/:id", userHandler.GetUser)                       // get user by id
	api.Post("/users/daftar", userHandler.RegisterUser)              // regiser user
	api.Get("/users/saldo/:no_rekening", userHandler.GetUserBalance) // get user current balance

	// user balance histories
	api.Post("/user-balance-histories/tabung", userBalanceHistoryHandler.AddBalance)     // add funds to user balance
	api.Post("/user-balance-histories/tarik", userBalanceHistoryHandler.WithdrawBalance) // withdraw fund from user balance
}
