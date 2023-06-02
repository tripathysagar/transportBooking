package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tripathysagar/transport-booking/db"
	"github.com/tripathysagar/transport-booking/types"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	fmt.Printf("%+v", userStore)

	return &UserHandler{userStore: userStore}
}

func HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	fmt.Printf("Trying to fetch data for userid : %v", id)
	return nil
}

func HandleGetUsers(c *fiber.Ctx) error {

	fmt.Printf("Trying to fetch all user data:")
	return nil
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var reqParams types.CreateUserParams

	if err := c.BodyParser(&reqParams); err != nil {
		return err
	}
	fmt.Println(reqParams)
	errMap := types.ValidateUser(reqParams)
	if len(errMap) != 0 {
		return c.JSON(errMap)
	}

	user, err := types.GetUserParams(reqParams)

	if err != nil {
		return err
	}
	user, err = h.userStore.PostUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(user)

}
