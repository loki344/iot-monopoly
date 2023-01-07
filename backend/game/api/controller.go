package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	service "iot-monopoly/game/adapter"
	"iot-monopoly/game/adapter/repository"
	"iot-monopoly/game/api/dto"
	"strconv"
)

func Routes(app *fiber.App) {

	//TODO look at how to organize routes https://github.com/gofiber/recipes/blob/2317ef83e51c79def9b5cb6adbfef5136f706f98/gorm-postgres/routes/routes.go

	app.Post("/games", func(c *fiber.Ctx) error {

		game := new(dto.GameDTO)

		if err := c.BodyParser(game); err != nil {
			fmt.Println("error = ", err)
			return fiber.ErrBadRequest
		}

		service.NewGame(game.PlayerCount)
		return c.Status(201).JSON(dto.GameDTOFromGame(service.GetCurrentGame()))
	})

	app.Patch("/games/current", func(c *fiber.Ctx) error {

		game := new(dto.GameDTO)

		if err := c.BodyParser(game); err != nil {
			fmt.Println("error = ", err)
			return fiber.ErrBadRequest
		}

		currentGame := service.GetCurrentGame()
		if game.Ended {
			currentGame.End("")
		}
		return c.Status(200).JSON(currentGame)
	})

	app.Get("/games/current", func(c *fiber.Ctx) error {

		return c.Status(200).JSON(dto.GameDTOFromGame(service.GetCurrentGame()))
	})

	app.Patch("/games/current/players/:id", func(c *fiber.Ctx) error {

		patchedPlayer := new(dto.PlayerDTO)

		if err := c.BodyParser(patchedPlayer); err != nil {
			fmt.Println("error = ", err)
			return fiber.ErrBadRequest
		}

		playerId := c.Params("id")
		currentGame := service.GetCurrentGame()
		currentGame.MovePlayer(playerId, patchedPlayer.Position)

		return c.Status(200).JSON(service.FindPlayerById(playerId))
	})

	app.Patch("/games/current/properties/:index", func(c *fiber.Ctx) error {

		patchRequest := new(dto.PropertyDTO)

		if err := c.BodyParser(patchRequest); err != nil {
			fmt.Println("error = ", err)
			return fiber.ErrBadRequest
		}

		param := c.Params("index")
		propertyIndex, _ := strconv.Atoi(param)

		service.BuyProperty(propertyIndex, patchRequest.OwnerId)

		return c.Status(200).JSON(dto.PropertyDTOFromProperty(repository.FindPropertyByIndex(propertyIndex)))
	})

	app.Patch("games/current/transactions/current", func(c *fiber.Ctx) error {

		transactionToPatch := new(dto.TransactionDTO)

		if err := c.BodyParser(transactionToPatch); err != nil {
			fmt.Println("error = ", err)
			return fiber.ErrBadRequest
		}

		if transactionToPatch.Accepted && transactionToPatch.SenderId != "" {
			transaction, err := service.ResolveCurrentTransaction(transactionToPatch.SenderId)
			if err != nil {
				return err
			}
			return c.Status(200).JSON(dto.TransactionDTOFromTransaction(transaction))
		}

		return c.SendStatus(400)
	})
	app.Patch("/games/current/cards/current", func(c *fiber.Ctx) error {

		cardDTO := new(dto.CardDTO)

		if err := c.BodyParser(cardDTO); err != nil {
			fmt.Println("error = ", err)
			return fiber.ErrBadRequest
		}

		if !cardDTO.Accepted {
			fmt.Println("Not confirmed, not doing anything")
			return c.SendStatus(400)
		}

		service.ConfirmCurrentCard()

		return c.SendStatus(200)
	})
}
