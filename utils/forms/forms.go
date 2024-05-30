package forms

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"reflect"
)

type ValidationError fiber.Map

func (e ValidationError) Error() string {
	return "ValidationError"
}

var validate = validator.New()

func BindAndValidate(c *fiber.Ctx, v interface{}) error {
	if err := c.BodyParser(v); err != nil {
		return ValidationError{
			"default": err.Error(),
		}
	}

	if err := validate.Struct(v); err != nil {
		errs := ValidationError{}
		ref := reflect.TypeOf(v).Elem()
		for _, e := range err.(validator.ValidationErrors) {
			if field, ok := ref.FieldByName(e.Field()); ok {
				k := field.Tag.Get("form")
				if len(k) == 0 {
					k = e.Field()
				}
				v := field.Tag.Get("errors." + e.ActualTag())
				if len(v) == 0 {
					v = e.Error()
				}
				errs[k] = v
			}
		}
		return errs
	}

	return nil
}
