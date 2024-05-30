package account

import (
	"github.com/gofiber/fiber/v2"
	"scaffold/utils/forms"
)

type LoginForm struct {
	Email    string `form:"email" validate:"required,email" errors.required:"Email address is required" errors.email:"Invalid email address"`
	Password string `form:"password" validate:"required,min=6" errors.required:"Password is required" errors.min:"Your password must be at least 6 characters"`
}

func FormValidation(c *fiber.Ctx) error {
	form := LoginForm{}
	if err := forms.BindAndValidate(c, &form); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"errors": err.(forms.ValidationError),
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"ok": true,
	})
}

func LoginPart(c *fiber.Ctx) error {
	return c.Render("partials/login", fiber.Map{
		"form":   LoginForm{},
		"errors": fiber.Map{},
	})
}

func Login(c *fiber.Ctx) error {
	form := LoginForm{}
	if err := forms.BindAndValidate(c, &form); err != nil {
		return c.Status(422).Render("partials/login", fiber.Map{
			"form":   form,
			"errors": err.(forms.ValidationError),
		})
	}
	return c.Status(202).SendString("ok")
}
