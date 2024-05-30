package main

import (
	"errors"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/django/v3"
	"log"
	"os"
	"os/signal"
	"scaffold/pkg/admin"
	"scaffold/pkg/front"
	"scaffold/utils/db"
	"scaffold/utils/tracker"
	"syscall"
)

const createSchema = "CREATE TABLE IF NOT EXISTS logs (id integer primary key, username text not null, created_at integer default (unixepoch()))"

func main() {
	db.New()
	defer db.Close()
	db.DB().MustExec(createSchema)

	app := fiber.New(fiber.Config{
		JSONEncoder:    json.Marshal,
		JSONDecoder:    json.Unmarshal,
		TrustedProxies: []string{},
		Views:          django.New("./views", ".html"),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			var e *fiber.Error
			if errors.As(err, &e) && e.Code < 500 {
				return c.Status(e.Code).SendString(e.Error())
			}
			log.Println(err)
			return c.SendStatus(500)
		},
	})
	app.Use(recover.New())
	app.Use(tracker.New())
	app.Static("/", "./public")
	app.Use(csrf.New(csrf.Config{KeyLookup: "cookie:csrf_"}))

	app.Mount("/admin", admin.Router())
	app.Mount("/", front.Router())

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
		<-ch
		app.Shutdown()
	}()

	if err := app.Listen(":3000"); err != nil {
		log.Println(err)
	}
}
