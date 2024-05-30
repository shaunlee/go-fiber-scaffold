package account

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
	"scaffold/utils/forms"
	"time"
)

type LoginForm struct {
	Email    string `form:"email" validate:"required,email" errors.required:"Email address is required" errors.email:"Invalid email address"`
	Password string `form:"password" validate:"required,min=6" errors.required:"Password is required" errors.min:"Your password must be at least 6 characters"`
	Pin      string `form:"pin" validate:"required,len=6" errors.required:"PIN is required" errors.len:"Invalid PIN"`
}

func LoginPart(c *fiber.Ctx) error {
	return c.Render("admin/partials/login", fiber.Map{
		"form":   LoginForm{},
		"errors": fiber.Map{},
	})
}

func Login(c *fiber.Ctx) error {
	form := LoginForm{}
	if err := forms.BindAndValidate(c, &form); err != nil {
		return c.Status(400).Render("admin/partials/login", fiber.Map{
			"form":   form,
			"errors": err.(forms.ValidationError),
		})
	}
	if token, err := login(&form); err != nil {
		log.Println(err)
		var e *fiber.Error
		if errors.As(err, &e) {
			return c.Status(e.Code).Render("admin/partials/login", fiber.Map{
				"form": form,
				"errors": fiber.Map{
					"default": e.Error(),
				},
			})
		}
		return err
	} else {
		c.Cookie(&fiber.Cookie{
			Name:     "sid",
			Value:    token,
			Path:     "/admin",
			Expires:  time.Now().Add(time.Hour),
			HTTPOnly: true,
			SameSite: "Strict",
		})
		c.Set("HX-Refresh", "true")
		return c.SendStatus(204)
	}
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "sid",
		Value:    "",
		Path:     "/admin",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		SameSite: "Strict",
	})
	return c.Redirect("/admin")
}
