package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	service "iot-monopoly/game/adapter"
)

func Routes(app *fiber.App) {

	//TODO look at how to organize routes https://github.com/gofiber/recipes/blob/2317ef83e51c79def9b5cb6adbfef5136f706f98/gorm-postgres/routes/routes.go

	app.Post("/games", func(c *fiber.Ctx) error {

		gameRequest := new(GameRequest)

		if err := c.BodyParser(gameRequest); err != nil {
			fmt.Println("error = ", err)
			return fiber.ErrBadRequest
		}

		service.StartGame(gameRequest.PlayerCount)
		return c.SendStatus(201)
	})

	app.Patch("/games/current", func(c *fiber.Ctx) error {

		gameEndRequest := new(GameEndRequest)

		if err := c.BodyParser(gameEndRequest); err != nil {
			fmt.Println("error = ", err)
			return fiber.ErrBadRequest
		}

		service.EndGame(gameEndRequest.Status)
		return c.SendStatus(200)
	})
}

type GameRequest struct {
	PlayerCount int `json:"playerCount"`
}

type GameEndRequest struct {
	Status string `json:"status"`
}
