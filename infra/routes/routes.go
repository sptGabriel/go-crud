package routes

import "github.com/gofiber/fiber/v2"

type Routes interface {
	Init(app *fiber.App)
}
