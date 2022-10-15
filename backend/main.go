package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"iot-monopoly/board/api"
	"iot-monopoly/finance"
	financeApi "iot-monopoly/finance/api"
)

func Init() {
	finance.StartEventHandler()
}

//start with CompileDaemon -command="./iot-monopoly"
func main() {

	Init()
	app := fiber.New()
	// Default config
	app.Use(cors.New())

	financeApi.Routes(app)
	movementApi.Routes(app)

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println(err)
		return
	}

}
