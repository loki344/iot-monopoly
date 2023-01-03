package financeApi

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	service "iot-monopoly/finance/adapter"
)

func Routes(app *fiber.App) {

	app.Get("/accounts", func(c *fiber.Ctx) error {

		return c.Status(200).JSON(service.GetAccounts())
	})

	app.Get("/accounts/:id", func(c *fiber.Ctx) error {

		account, err := service.GetAccountById(c.Params("id"))

		if err != nil {
			return c.Status(404).JSON(err)
		}

		return c.Status(200).JSON(account)
	})

	app.Patch("/transactions/latest", func(c *fiber.Ctx) error {

		transactionToPatch := new(TransactionAcceptedPatch)

		if err := c.BodyParser(transactionToPatch); err != nil {
			fmt.Println("error = ", err)
			return fiber.ErrBadRequest
		}

		if transactionToPatch.Accepted && transactionToPatch.CardId != "" {
			err := service.ResolveLatestTransaction(transactionToPatch.CardId)
			if err != nil {
				return err
			}
			return c.SendStatus(200)
		}

		return c.SendStatus(400)
	})

}

type TransactionAcceptedPatch struct {
	Accepted bool   `json:"accepted"`
	CardId   string `json:"cardId"`
}
