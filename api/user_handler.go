package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tripathysagar/transport-booking/types"
)

func HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	fmt.Printf("Trying to fetch data for userid : %v", id)
	return nil
}

func HandlePostUser(c *fiber.Ctx) error {
	var reqParams types.CreateUserParams

	if err := c.BodyParser(&reqParams); err != nil {
		return err
	}
	fmt.Println(reqParams)
	return c.JSON(reqParams)

}
