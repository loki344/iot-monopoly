package financeApi

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"iot-monopoly/finance"
	"iot-monopoly/finance/domain"
)

func Routes(app *fiber.App) {

	app.Post("/transactions", func(c *fiber.Ctx) error {

		newTransaction := new(financeDomain.Transaction)

		if err := c.BodyParser(newTransaction); err != nil {
			fmt.Println("error = ", err)
			return fiber.ErrBadRequest
		}

		newTransaction = financeDomain.NewTransactionFromTransactionDTO(newTransaction)

		_, err := finance.AddTransaction(*newTransaction)
		if err != nil {
			return fiber.ErrBadRequest
		}

		return c.Status(200).JSON(newTransaction)
	})

	app.Patch("/transactions/:id", func(c *fiber.Ctx) error {

		transactionToPatch := new(financeDomain.Transaction)

		if err := c.BodyParser(transactionToPatch); err != nil {
			fmt.Println("error = ", err)
			return fiber.ErrBadRequest
		}

		err := validatePatchTransaction(*transactionToPatch, finance.GetTransaction(c.Params("id")))
		if err != nil {
			return err
		}

		if transactionToPatch.Accepted {
			finance.ResolveTransaction(transactionToPatch.Id)
		}
		return c.Status(200).JSON(transactionToPatch)
	})

}

func validatePatchTransaction(toCheck financeDomain.Transaction, transaction financeDomain.Transaction) error {

	if toCheck.Id != transaction.Id || toCheck.SenderId != transaction.SenderId || toCheck.RecipientId != transaction.RecipientId || toCheck.Amount != transaction.Amount {
		return errors.New("Transaction invalid, only changing the accept state of a transaction is allowed")
	}

	return nil
}
