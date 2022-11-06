package eventing

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/vmware/transport-go/model"
	"log"
)

var EXTERNAL_CHANNELS = [2]ChannelName{PROPERTY_BUY_QUESTION, TRANSACTION_REQUESTED}

func StartWebsocket(app *fiber.App) {

	registerWebsocket(app)
}

func registerWebsocket(app *fiber.App) fiber.Router {
	return app.Get("/ws", websocket.New(func(c *websocket.Conn) {

		for i := range EXTERNAL_CHANNELS {

			eventHandler := ListenRequestStream(EXTERNAL_CHANNELS[i])
			eventHandler.Handle(
				func(eventMessage *model.Message) {

					err := c.WriteJSON(eventMessage.Payload)
					if err != nil {
						fmt.Println(err)
						return
					}
				},
				func(err error) {
					fmt.Println(err)
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
