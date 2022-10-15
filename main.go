package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"iot-monopoly/board"
	"iot-monopoly/board/api"
	"iot-monopoly/board/domain"
	"iot-monopoly/finance"
	financeApi "iot-monopoly/finance/api"
)

func Init() {
	finance.StartEventHandler()
}

//start with CompileDaemon -command="./iot-monopoly"
func main() {

	Init()
	//TODO get player infos from board
	playerOneId := uuid.New().String()
	playerTwoId := uuid.New().String()
	board.StartGame([]boardDomain.Player{{playerOneId, 0, 1000}, {playerTwoId, 0, 1000}})

	app := fiber.New()

	financeApi.Routes(app)
	movementApi.Routes(app)

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println(err)
		return
	}

}
