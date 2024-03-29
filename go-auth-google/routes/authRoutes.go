package routes

import (
	"encoding/json"
	"go-auth-google/backend/auth"
	config "go-auth-google/backend/config"
	"go-auth-google/backend/errset"
	"go-auth-google/backend/model"
	"go-auth-google/backend/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func HomeRoute(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func LoginRoute(c *fiber.Ctx) error {
	url := config.GOOGLEAuthConfig.AuthCodeURL(config.State)
	return c.Redirect(url, http.StatusTemporaryRedirect)
}

func AuthCallBack(c *fiber.Ctx) error {

	content, err := auth.GetUserInfo(c.FormValue("state"), c.FormValue("code"))
	if err != nil {
		return errset.ReturnError(err, c)
	}

	var user model.User
	json.Unmarshal(content, &user)
	token, err := service.GenerateJwtTokenService(&user)
	if err != nil {
		return errset.ReturnError(err, c)
	}

	return c.Status(http.StatusOK).SendString(token)
}
