package movementApi

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"iot-monopoly/board"
	boardDomain "iot-monopoly/board/domain"
)

func Routes(app *fiber.App) {

	//TODO look at how to organize routes https://github.com/gofiber/recipes/blob/2317ef83e51c79def9b5cb6adbfef5136f706f98/gorm-postgres/routes/routes.go

	app.Get("/players", func(c *fiber.Ctx) error {

		return c.Status(200).JSON(board.GetPlayers())
	})

	app.Patch("/players/:id", func(c *fiber.Ctx) error {

		player := new(boardDomain.Player)

		if err := c.BodyParser(player); err != nil {
			fmt.Println("error = ", err)
			return fiber.ErrBadRequest
		}

		playerId := c.Params("id")

		board.MovePlayer(playerId, player.Position)

		return c.Status(201).JSON(board.GetPlayer(playerId))
	})

	app.Patch("/fields/:id", func(c *fiber.Ctx) error {

		patchRequest := new(PropertyPatchRequest)

		if err := c.BodyParser(patchRequest); err != nil {
			fmt.Println("error = ", err)
			return fiber.ErrBadRequest
		}

		propertyId := c.Params("id")

		transactionId := board.BuyProperty(propertyId, patchRequest.OwnerId)

		return c.Status(200).JSON(transactionId)
	})

	app.Post("/games", func(c *fiber.Ctx) error {

		gameRequest := new(GameRequest)

		if err := c.BodyParser(gameRequest); err != nil {
			fmt.Println("error = ", err)
			return fiber.ErrBadRequest
		}

		players, err := board.StartGame(gameRequest.PlayerCount)
		if err != nil {
			return err
		}

		return c.Status(201).JSON(players)
	})

}

type GameRequest struct {
	PlayerCount int `json:"playerCount"`
}

type PropertyPatchRequest struct {
	Id      string `json:"id"`
	OwnerId string `json:"ownerId"`
}
