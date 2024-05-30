package front

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"scaffold/utils"
)

var LANGS = []string{"en", "zh-Hant", "zh", "ja", "ko"}

func langRedirect(c *fiber.Ctx) error {
	lang := c.AcceptsLanguages(LANGS...)
	if len(lang) == 0 {
		lang = "en"
	}
	return c.Redirect("/" + lang)
}

func indexPage(c *fiber.Ctx) error {
	dblog, err := writeLog("Shaun")
	if err != nil {
		return fiber.NewError(500, fmt.Sprintf("Failed writeLog: %s", err.Error()))
	}
	return c.Render("index", fiber.Map{
		"is_prod": true,
		"name":    "Shaun",
		"log":     dblog,
	})
}

func jsonData(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"ok": true,
	})
}

func randomCode(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"code": utils.RandomCode(24),
	})
}
