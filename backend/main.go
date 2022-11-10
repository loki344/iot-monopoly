package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"iot-monopoly/board/api"
	"iot-monopoly/eventing"
	"iot-monopoly/eventing/config"
	"iot-monopoly/finance"
	financeApi "iot-monopoly/finance/api"
)

func Init() {
	config.Init()
	finance.StartEventListeners()
}

//start with CompileDaemon -command="./iot-monopoly"
func main() {

	Init()
	app := fiber.New()
	// Default config
	app.Use(cors.New())
	app.Use(logger.New())

	eventing.StartWebsocket(app)
	financeApi.Routes(app)
	movementApi.Routes(app)

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println(err)
		return
	}

}
