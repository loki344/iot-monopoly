package financeApi

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"iot-monopoly/finance"
	"iot-monopoly/finance/financeDomain"
)

func Routes(app *fiber.App) {

	app.Patch("/transactions/:id", func(c *fiber.Ctx) error {

		transactionToPatch := new(financeDomain.Transaction)

		if err := c.BodyParser(transactionToPatch); err != nil {
			fmt.Println("error = ", err)
			return fiber.ErrBadRequest
		}

		err := validateTransaction(*transactionToPatch, finance.GetTransaction(c.Params("id")))
		if err != nil {
			return err
		}

		if transactionToPatch.Accepted {
			finance.ResolveTransaction(transactionToPatch.Id())
		}
		return c.Status(200).JSON(transactionToPatch)
	})

}

func validateTransaction(toCheck financeDomain.Transaction, transaction financeDomain.Transaction) error {

	if toCheck.Id() != transaction.Id() || toCheck.SenderId() != transaction.SenderId() || toCheck.RecipientId() != transaction.RecipientId() || toCheck.Amount() != transaction.Amount() {
		return errors.New("Transaction invalid, only changing the accept state of a transaction is allowed")
	}

	return nil
}
