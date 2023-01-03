package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"iot-monopoly/communication"
	"iot-monopoly/communication/config"
	financeAdapter "iot-monopoly/finance/adapter"
	financeApi "iot-monopoly/finance/api"
	gameApi "iot-monopoly/game/api"
	gameEventsAdapter "iot-monopoly/gameEvents/adapter"
	gameEventsApi "iot-monopoly/gameEvents/api"
	playerAdapter "iot-monopoly/player/adapter"
	"iot-monopoly/player/api"
	propertyAdapter "iot-monopoly/property/adapter"
	propertyApi "iot-monopoly/property/api"
)

func Init() {
	config.Init()
	playerAdapter.StartEventListeners()
	financeAdapter.StartEventListeners()
	propertyAdapter.StartEventListeners()
	gameEventsAdapter.StartEventListeners()
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
