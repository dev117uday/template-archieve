package server

import (
	"go-postgres/conf"
	"go-postgres/logrus"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func StartServer() {

	app := fiber.New()
	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: conf.FRONTEND_URL,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/user", AuthUser)
	app.Get("/user/:id", GetUser)

	err := app.Listen(conf.PORT)
	if err != nil {
		logrus.LogIt(0, "StartServer", "unable to start server", err)
	}
}
