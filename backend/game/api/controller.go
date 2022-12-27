package gameApi

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"iot-monopoly/game"
)

func Routes(app *fiber.App) {

	//TODO look at how to organize routes https://github.com/gofiber/recipes/blob/2317ef83e51c79def9b5cb6adbfef5136f706f98/gorm-postgres/routes/routes.go

	app.Post("/games", func(c *fiber.Ctx) error {

		gameRequest := new(GameRequest)

		if err := c.BodyParser(gameRequest); err != nil {
			fmt.Println("error = ", err)
			return fiber.ErrBadRequest
		}

		game.StartGame(gameRequest.PlayerCount)
		return c.SendStatus(201)
	})

}

type GameRequest struct {
	PlayerCount int `json:"playerCount"`
}
