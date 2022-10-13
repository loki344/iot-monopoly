package financeApi

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"iot-monopoly/board"
	"iot-monopoly/finance"
	"iot-monopoly/finance/financeDomain"
)

func Routes(app *fiber.App) {

	//TODO look at how to organize routes https://github.com/gofiber/recipes/blob/2317ef83e51c79def9b5cb6adbfef5136f706f98/gorm-postgres/routes/routes.go
	app.Post("/transactions", func(c *fiber.Ctx) error {

		transaction := new(financeDomain.Transaction)

		if err := c.BodyParser(transaction); err != nil {
			fmt.Println("error = ", err)
			return fiber.ErrBadRequest
		}

		sender := board.GetPlayer(transaction.SenderId())
		if sender.Balance < transaction.Amount() {
			err := fmt.Sprintf("Sender %s has insufficient balance. Balance: %d, Amount: %d", sender.Id, sender.Balance, transaction.Amount())
			return fiber.NewError(500, err)
		} else {
			transaction := *financeDomain.NewTransaction(transaction.Id(), transaction.RecipientId(), transaction.SenderId(), transaction.Amount())
			finance.AddTransaction(transaction)
		}

		return c.Status(201).JSON(transaction)
	})

	app.Patch("/transactions/:id", func(c *fiber.Ctx) error {

		transactionToPatch := new(financeDomain.Transaction)

		if err := c.BodyParser(transactionToPatch); err != nil {
			fmt.Println("error = ", err)
			return fiber.ErrBadRequest
		}

		checkIntegrity(transactionToPatch, finance.GetTransaction(c.Params("id")))

		if transactionToPatch.Accepted {
			transactionToPatch.Resolve()
		}
		return c.Status(200).JSON(transactionToPatch)
	})

}

func checkIntegrity(toCheck *financeDomain.Transaction, transaction *financeDomain.Transaction) error {

	if toCheck.Id() != transaction.Id() || toCheck.SenderId() != transaction.SenderId() || toCheck.RecipientId() != transaction.RecipientId() || toCheck.Amount() != transaction.Amount() {
		return errors.New("Transaction invalid, only changing the accept state of a transaction is allowed")
	}

	return nil
}
