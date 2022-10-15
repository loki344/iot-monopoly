package movementApi

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"iot-monopoly/board"
	boardDomain "iot-monopoly/board/domain"
)

func Routes(app *fiber.App) {

	//TODO look at how to organize routes https://github.com/gofiber/recipes/blob/2317ef83e51c79def9b5cb6adbfef5136f706f98/gorm-postgres/routes/routes.go
	app.Patch("/players/:id", func(c *fiber.Ctx) error {

		player := new(boardDomain.Player)

		if err := c.BodyParser(player); err != nil {
			fmt.Println("error = ", err)
			return fiber.ErrBadRequest
		}

		err := validateMessage(player, board.GetPlayer(c.Params("id")))
		if err != nil {
			return err
		}

		board.MovePlayer(player.Id, player.Position)

		return c.Status(201).JSON(player)
	})

}

func validateMessage(toCheck *boardDomain.Player, player *boardDomain.Player) error {

	if toCheck.Id != player.Id || toCheck.Balance != player.Balance {
		return errors.New("Patch player invalid, only changing the position field of a player is allowed")
	}

	return nil
}
