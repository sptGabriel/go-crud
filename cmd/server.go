package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sptGabriel/go-ddd/infra/routes"
	"github.com/sptGabriel/go-ddd/presentation/controllers"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Live"})
	})
	personController := controllers.NewPersonController("a")
	personRouter := routes.NewPersonRoutes(personController)
	personRouter.Init(app)
	err := app.Listen(":8080")
	fmt.Println(err)
}
