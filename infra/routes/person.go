package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sptGabriel/go-ddd/presentation/controllers"
)

type personRoutes struct {
	personController *controllers.PersonController
}

func NewPersonRoutes(personController *controllers.PersonController) Routes {
	return &personRoutes{personController}
}

func (route *personRoutes) Init(app *fiber.App) {
	app.Get("/person/:personId", route.personController.GetPerson)
	app.Post("/person", route.personController.CreateNewPerson)
}
