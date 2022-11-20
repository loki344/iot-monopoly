package eventing

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/mustafaturan/bus/v3"
)

var EXTERNAL_CHANNELS = []ChannelName{PROPERTY_BUY_QUESTION, TRANSACTION_REQUEST, TRANSACTION_RESOLVED}

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

			RegisterEventHandler(bus.Handler{
				Handle: func(ctx context.Context, e bus.Event) {
					// TODO this is a workaround, it would be better to deregsiter the eventhandler when the websocket connection is closed
					if c.Conn == nil {
						return
					}
					err := c.WriteJSON(e.Data)
					if err != nil {
						fmt.Println(err)
						return
					}
				},
				Matcher: string(EXTERNAL_CHANNELS[i]),
			})
		}
	}))
}
