package account

import (
	"github.com/gofiber/fiber/v2"
)

func login(form *LoginForm) (string, error) {
	if form.Password != "123456" {
		return "", fiber.NewError(400, "Incorrect email address or password")
	}

	return "ok", nil
}
