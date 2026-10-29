package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Connect Database
	db, error := ConnectDatabase()

	if error != nil {
		fmt.Println(error, db)
	}

	// Initialize Fiber app
	app := fiber.New()

	studentRoute(app, db)

	// Start server
	if error := app.Listen(":3030"); error != nil {
		fmt.Println(error)
	}
}
