package main

import (
	"fmt"
	"github.com/miceremwirigi/go-sales-api/config"
	
)
git 
func main() {

	fmt.Println("Go sales api started")
	config.Connect()
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
