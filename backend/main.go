package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"iot-monopoly/communication"
	"iot-monopoly/communication/config"
	"iot-monopoly/finance"
	financeApi "iot-monopoly/finance/api"
	gameApi "iot-monopoly/game/api"
	"iot-monopoly/gameEvents"
	gameEventsApi "iot-monopoly/gameEvents/api"
	"iot-monopoly/player"
	"iot-monopoly/player/api"
	"iot-monopoly/property"
	propertyApi "iot-monopoly/property/api"
)

func Init() {
	config.Init()
	player.StartEventListeners()
	finance.StartEventListeners()
	property.StartEventListeners()
	gameEvents.StartEventListeners()
}

//start with CompileDaemon -command="./iot-monopoly"
func main() {

	Init()
	app := fiber.New()
	// Default config
	app.Use(cors.New())
	app.Use(logger.New())

	communication.StartWebsocket(app)
	gameApi.Routes(app)
	gameEventsApi.Routes(app)
	financeApi.Routes(app)
	playerApi.Routes(app)
	propertyApi.Routes(app)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println(err)
		return
	}

}
