package propertyApi

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"iot-monopoly/property"
)

func Routes(app *fiber.App) {

	//TODO look at how to organize routes https://github.com/gofiber/recipes/blob/2317ef83e51c79def9b5cb6adbfef5136f706f98/gorm-postgres/routes/routes.go

	app.Patch("/properties/:id", func(c *fiber.Ctx) error {

		patchRequest := new(PropertyPatchRequest)

		if err := c.BodyParser(patchRequest); err != nil {
			fmt.Println("error = ", err)
			return fiber.ErrBadRequest
		}

		propertyId := c.Params("id")

		transactionId := property.BuyProperty(propertyId, patchRequest.OwnerId)

		return c.Status(200).JSON(transactionId)
	})

}

type PropertyPatchRequest struct {
	OwnerId string `json:"ownerId"`
}
