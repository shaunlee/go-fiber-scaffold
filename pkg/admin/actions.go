package admin

import (
	"github.com/gofiber/fiber/v2"
)

func dashboard(c *fiber.Ctx) error {
	return c.Render("admin/dashboard", fiber.Map{})
}
