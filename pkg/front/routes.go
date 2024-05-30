package front

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"scaffold/pkg/front/account"
	"strings"
)

func Router() *fiber.App {
	var app = fiber.New()

	app.Get("/", langRedirect)

	app.Get("/check", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	app.Get("/error", func(c *fiber.Ctx) error {
		return c.Status(500).SendString("Something wrong")
	})

	app.Get("/no-content", func(c *fiber.Ctx) error {
		return c.SendStatus(204)
	})

	app.Get("/json", jsonData)
	app.Post("/json", account.FormValidation)

	app.Get("/api/random", randomCode)

	pgs := app.Group("/:lang<regex(^("+strings.Join(LANGS, "|")+"))>", func(c *fiber.Ctx) error {
		if !lo.Contains(LANGS, c.Params("lang")) {
			return c.SendStatus(404)
		}
		c.Bind(fiber.Map{
			"lang": c.Params("lang"),
		})
		return c.Next()
	})

	pgs.Get("/", indexPage)

	pgs.Get("/login", account.LoginPart)
	pgs.Post("/login", account.Login)

	return app
}
