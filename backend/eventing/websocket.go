package eventing

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/mustafaturan/bus/v3"
	"log"
)

var EXTERNAL_CHANNELS = [2]ChannelName{PROPERTY_BUY_QUESTION, TRANSACTION_REQUEST}

func StartWebsocket(app *fiber.App) {

	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	registerWebsocket(app)
}

func registerWebsocket(app *fiber.App) fiber.Router {
	return app.Get("/ws", websocket.New(func(c *websocket.Conn) {

		for i := range EXTERNAL_CHANNELS {

			//TODO we have an issue that crashes when we refresh the browser and then send events
			RegisterEventHandler(bus.Handler{
				Handle: func(ctx context.Context, e bus.Event) {
					err := c.WriteJSON(e.Data)
					if err != nil {
						fmt.Println(err)
						return
					}
				},
				Matcher: string(EXTERNAL_CHANNELS[i]),
			})
		}

		var (
			msg []byte
			err error
		)
		for {
			if _, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)
		}

	}))
}
