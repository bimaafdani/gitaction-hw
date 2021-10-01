package main

import (
	"gitaction_hw/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/follower/:username", handler.GetAll)
	app.Get("/:userid/detail", handler.GetByUserId)
	app.Listen(":3000")
}
