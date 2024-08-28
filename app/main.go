package main

import (
	"fmt"

	"github.com/Mayank-Sharma-27/KeepInMindBackend.git/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "KeepInMind",
		ServerHeader: "Fiber",
	})
	server := app.Group("/api")

	privateRoutes := server.Use()
	handlers.NewUserHandler(privateRoutes.Group("/user"))
	fmt.Print("server started")
	app.Listen(fmt.Sprintf(":8080"))
}
