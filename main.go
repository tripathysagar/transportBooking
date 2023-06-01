package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/tripathysagar/transport-booking/api"
)

var errConf = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	}}

func main() {
	// Fiber instance
	app := fiber.New(errConf)
	apiv1 := app.Group("/app/v1")

	// Routes
	apiv1.Get("/users", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)
	apiv1.Post("/users", api.HandlePostUser)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
