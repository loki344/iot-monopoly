package financeApi

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"iot-monopoly/finance"
)

func Routes(app *fiber.App) {

	app.Get("/accounts", func(c *fiber.Ctx) error {

		return c.Status(200).JSON(finance.GetAccounts())
	})

	app.Get("/accounts/:id", func(c *fiber.Ctx) error {

		account, err := finance.GetAccountById(c.Params("id"))

		if err != nil {
			return c.Status(404).JSON(err)
		}

		return c.Status(200).JSON(account)
	})

	app.Patch("/transactions/:id", func(c *fiber.Ctx) error {

		transactionToPatch := new(TransactionAcceptedPatch)

		if err := c.BodyParser(transactionToPatch); err != nil {
			fmt.Println("error = ", err)
			return fiber.ErrBadRequest
		}

		if transactionToPatch.Accepted {
			finance.ResolveTransaction(c.Params("id"))
		}
		return c.Status(200).JSON(finance.GetTransaction(c.Params("id")))
	})

}

type TransactionAcceptedPatch struct {
	Accepted bool `json:"accepted"`
}
