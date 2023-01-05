package eventing

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/mustafaturan/bus/v3"
)

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

var EXTERNAL_CHANNELS = []ChannelName{PLAYER_ON_UNOWNED_FIELD, TRANSACTION_CREATED, TRANSACTION_RESOLVED, CARD_DREW, LAP_FINISHED}

func registerWebsocket(app *fiber.App) fiber.Router {
	return app.Get("/ws", websocket.New(func(c *websocket.Conn) {

		for i := range EXTERNAL_CHANNELS {

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

	}))
}
