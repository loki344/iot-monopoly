package playerApi

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"iot-monopoly/player"
)

func Routes(app *fiber.App) {

	//TODO look at how to organize routes https://github.com/gofiber/recipes/blob/2317ef83e51c79def9b5cb6adbfef5136f706f98/gorm-postgres/routes/routes.go

	type PlayerResponse struct {
		Players         []*player.PlayerDTO `json:"players"`
		CurrentPlayerId string              `json:"currentPlayerId"`
	}

	app.Get("/players", func(c *fiber.Ctx) error {

		return c.Status(200).JSON(PlayerResponse{Players: player.GetPlayers(), CurrentPlayerId: player.GetCurrentPlayer().Id})
	})

	app.Patch("/players/:id", func(c *fiber.Ctx) error {

		patchedPlayer := new(player.PlayerDTO)

		if err := c.BodyParser(patchedPlayer); err != nil {
			fmt.Println("error = ", err)
			return fiber.ErrBadRequest
		}

		playerId := c.Params("id")
		player.MovePlayer(playerId, patchedPlayer.Position)

		return c.Status(201).JSON(player.GetPlayer(playerId))
	})

}
