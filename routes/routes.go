package routes

import (
	"github.com/gofiber/fiber/v2"
	authControllers "github.com/miceremwirigi/go-sales-api/controllers/authControllers"
	controllers "github.com/miceremwirigi/go-sales-api/controllers/cashierControllers"
)

// setup routes
func Setup(app *fiber.App) {
	// cashier session routes
	app.Post("/cashiers/:cashierId/login", authControllers.Login)
	app.Get("/cashiers/:cashierId/logout", authControllers.Logout)
	app.Post("/cashiers/:cashierId/passcode", authControllers.Passcode)

	// cashier CRUD routes
	app.Post("/newcashier", controllers.CreateCashier)
	app.Get("/cashiers", controllers.CashiersList)
	app.Get("/cashiers/:cashierId/details", controllers.GetCashierDetails)
	app.Delete("/cashiers/:cashierId/delete", controllers.DeleteCashier)
	app.Put("/cashiers/:cashierId/update", controllers.UpdateCashier)
}
