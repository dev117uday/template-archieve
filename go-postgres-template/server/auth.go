package server

import (
	"encoding/json"
	"fmt"
	"go-postgres/db"
	"go-postgres/errset"
	"go-postgres/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func AuthUser(c *fiber.Ctx) error {

	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		fmt.Println(err)
	}

	err := db.AddUser(user)

	if err != nil {
		return errset.ReturnError(err, c)
	}

	return c.SendStatus(http.StatusOK)

}

func GetUser(c *fiber.Ctx) error {

	user, err := db.GetUser(c.Params("id"))
	if err != nil {
		return errset.ReturnError(err, c)
	}

	data, _ := json.Marshal(user)
	return c.Status(http.StatusOK).Send(data)

}
