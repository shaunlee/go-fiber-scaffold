package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"scaffold/pkg/admin/account"
)

func Router() *fiber.App {
	var app = fiber.New()

	var visibleUrls = []string{"/admin/login", "/admin/logout"}

	app.Use(func(c *fiber.Ctx) error {
		if lo.Contains(visibleUrls, c.OriginalURL()) {
			return c.Next()
		}
		if sid := c.Cookies("sid"); sid == "ok" {
			return c.Next()
		}
		return c.Render("admin/login", fiber.Map{
			"form":   account.LoginForm{},
			"errors": fiber.Map{},
		})
	})

	app.Get("/", dashboard)

	app.Get("/login", account.LoginPart)
	app.Post("/login", account.Login)
	app.Get("/logout", account.Logout)

	return app
}
