package movementApi

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"iot-monopoly/eventing"
	eventingDomain "iot-monopoly/eventing/domain"
)

var sensorEvents []eventingDomain.SensorEvent

func Routes(app *fiber.App) {

	//TODO look at how to organize routes https://github.com/gofiber/recipes/blob/2317ef83e51c79def9b5cb6adbfef5136f706f98/gorm-postgres/routes/routes.go
	app.Post("/sensor-events", func(c *fiber.Ctx) error {

		sensorEvent := new(eventingDomain.SensorEvent)

		if err := c.BodyParser(sensorEvent); err != nil {
			fmt.Println("error = ", err)
			return fiber.ErrBadRequest
		}

		sensorEvents = append(sensorEvents, *sensorEvent)
		eventing.FireEvent(eventing.MOVEMENT, sensorEvent)

		return c.Status(201).JSON(sensorEvent)
	})

}
