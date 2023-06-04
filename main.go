package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/tripathysagar/transport-booking/api"
	"github.com/tripathysagar/transport-booking/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DBNAME = "taxi_reservation"

var errConf = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	}}

func main() {
	// Fiber instance
	app := fiber.New(errConf)
	apiv1 := app.Group("/app/v1")

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.DBURI))

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()
	x := db.NewMongoStore(client, DBNAME)
	fmt.Printf("%+v", x)
	userHandler := api.NewUserHandler(x)
	// Routes
	apiv1.Get("/users", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	apiv1.Post("/user", userHandler.HandlePostUser)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
