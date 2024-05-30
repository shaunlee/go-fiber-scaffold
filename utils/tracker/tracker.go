package tracker

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strings"
	"time"
)

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		t := time.Now()

		id := fmt.Sprintf("%x", t.Nanosecond())
		ip := c.IP()
		method := c.Method()
		path := c.OriginalURL()
		msg := ""
		if c.Is("json") {
			if body := c.Body(); len(body) > 0 {
				msg = string(body)
			}
		}
		log.Println(id, ip, "<-", method, path, msg, c.Get("User-Agent"))

		err := c.Next()

		code := c.Response().Header.StatusCode()
		if err != nil {
			if ferr, ok := err.(*fiber.Error); ok {
				code = ferr.Code
			} else {
				code = 500
			}
		}

		rmsg := ""
		if strings.HasPrefix(c.GetRespHeader("Content-Type"), "application/json") {
			if body := c.Response().Body(); len(body) > 0 {
				rmsg = string(body)
			}
		}
		log.Println(id, ip, "->", method, code, path, time.Since(t), rmsg)

		return err
	}
}
