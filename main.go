package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	db "github.com/miceremwirigi/go-sales-api/config/dbConnection"
	routes "github.com/miceremwirigi/go-sales-api/routes"
)

func main() {

	fmt.Println("Go sales api started")
	db.Connect()

	// create new fiber app
	app := fiber.New()
	// middleware to match anything
	// app.Use(app)

	// import routes
	routes.Setup(app)

	// listen on port
	app.Listen(":30001")

}

// // Create instance of fiber
// app := fiber.New()

// // Create httphandler
// app.Get("/testApi", func(ctx *fiber.Ctx) error {
// 	return ctx.Status(200).JSON(fiber.Map{
// 		"success": true,
// 		"message": "Go fiber first api created successfully",
// 	})
// })

// // listen on port
// app.Listen(":3000")
