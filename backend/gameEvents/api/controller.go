package gameEventsApi

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	adapter "iot-monopoly/gameEvents/adapter"
)

func Routes(app *fiber.App) {

	//TODO look at how to organize routes https://github.com/gofiber/recipes/blob/2317ef83e51c79def9b5cb6adbfef5136f706f98/gorm-postgres/routes/routes.go

	app.Patch("/card-events/latest", func(c *fiber.Ctx) error {

		confirmedDTO := new(ConfirmedDTO)

		if err := c.BodyParser(confirmedDTO); err != nil {
			fmt.Println("error = ", err)
			return fiber.ErrBadRequest
		}

		if !confirmedDTO.Confirmed {
			fmt.Println("Not confirmed, not doing anything")
			return c.SendStatus(400)
		}

		adapter.ConfirmCard()

		return c.SendStatus(200)
	})
}

type ConfirmedDTO struct {
	Confirmed bool `json:"confirmed"`
}
