package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	service "iot-monopoly/player/adapter"
)

func Routes(app *fiber.App) {

	//TODO look at how to organize routes https://github.com/gofiber/recipes/blob/2317ef83e51c79def9b5cb6adbfef5136f706f98/gorm-postgres/routes/routes.go

	type PlayerResponseDTO struct {
		Players         []*service.PlayerDTO `json:"players"`
		CurrentPlayerId string               `json:"currentPlayerId"`
	}

	app.Get("/players", func(c *fiber.Ctx) error {

		return c.Status(200).JSON(PlayerResponseDTO{Players: service.GetPlayers(), CurrentPlayerId: service.GetCurrentPlayer().Id})
	})

	app.Patch("/players/:id", func(c *fiber.Ctx) error {

		patchedPlayer := new(service.PlayerDTO)

		if err := c.BodyParser(patchedPlayer); err != nil {
			fmt.Println("error = ", err)
			return fiber.ErrBadRequest
		}

		playerId := c.Params("id")
		service.MovePlayer(playerId, patchedPlayer.Position)

		return c.Status(201).JSON(service.GetPlayer(playerId))
	})

}
