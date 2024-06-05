package sse

import (
	"bufio"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/nats.go"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"github.com/valyala/fasthttp"
	"log"
	"strings"
	"time"
)

func ping(w *bufio.Writer) error {
	data, _ := sjson.Set("", "time", time.Now().Format(time.RFC3339))
	fmt.Fprintf(w, "event: ping\ndata: %s\n\n", data)
	return w.Flush()
}

func sendMessage(w *bufio.Writer, msg *nats.Msg) error {
	event := gjson.GetBytes(msg.Data, "event").String()
	data := gjson.GetBytes(msg.Data, "data").String()
	fmt.Fprintf(w, "event: %s\ndata: %s\n\n", event, data)
	return w.Flush()
}

type MsgChanSub struct {
	Chan chan *nats.Msg
	Sub  *nats.Subscription
}

func Router() *fiber.App {
	var app = fiber.New()

	app.Get("/:channels", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/event-stream; charset=utf-8")
		c.Set("Cache-Control", "no-cache")
		c.Set("Connection", "keep-alive")
		c.Set("Transfer-Encoding", "chunked")

		channels := strings.Split(c.Params("channels"), ",")

		exit := c.Context().Done()

		c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
			if err := ping(w); err != nil {
				return
			}

			log.Println("Welcome!")
			defer log.Println("Goodbye!")

			nc, _ := nats.Connect("nats://127.0.0.1:4222")
			defer nc.Drain()

			msgs := make(chan *nats.Msg)
			defer close(msgs)
			msgChans := make(map[string]MsgChanSub)
			for _, c := range channels {
				mc := make(chan *nats.Msg)
				sub, _ := nc.ChanSubscribe(c, mc)
				msgChans[c] = MsgChanSub{mc, sub}
				go func(mc chan *nats.Msg) {
					for ch := range mc {
						msgs <- ch
					}
				}(mc)
			}
			defer func() {
				for _, mc := range msgChans {
					close(mc.Chan)
					mc.Sub.Unsubscribe()
				}
			}()

			for loop := true; loop; {
				select {
				case msg := <-msgs:
					if err := sendMessage(w, msg); err != nil {
						loop = false
						break
					}
				case <-time.After(5 * time.Second):
					if err := ping(w); err != nil {
						loop = false
						break
					}
				case <-exit:
					loop = false
					break
				}
			}
		}))

		return nil
	})

	return app
}
