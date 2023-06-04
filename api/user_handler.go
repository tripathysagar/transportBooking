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

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println("Trying to fetch user datafor id : ", id)

	user, err := h.userStore.GetUserByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(user)

}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {

	fmt.Println("Trying to fetch all user data ")
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(users)

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

	fmt.Println("tying to post user : ", user)
	if err != nil {
		return err
	}
	user, err = h.userStore.PostUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(user)

}
