package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"iot-monopoly/eventing"
	"iot-monopoly/eventing/config"
	gameService "iot-monopoly/game/adapter"
	gameApi "iot-monopoly/game/api"
)

func Init() {
	config.Init()
	gameService.StartEventListeners()
}

//start with CompileDaemon -command="./iot-monopoly"
func main() {

	Init()
	app := fiber.New()
	// Default config
	app.Use(cors.New())
	app.Use(logger.New())

	eventing.StartWebsocket(app)
	gameApi.Routes(app)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println(err)
		return
	}

}
