package main

import (
	"log"
	"net/http"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/sptGabriel/go-ddd/application/handlers"
	person "github.com/sptGabriel/go-ddd/domain/person/commands"
	"github.com/sptGabriel/go-ddd/infra/commandBus"
	"github.com/sptGabriel/go-ddd/infra/database"
	repositoriesImpl "github.com/sptGabriel/go-ddd/infra/repositories"
	"github.com/sptGabriel/go-ddd/infra/routes"
	"github.com/sptGabriel/go-ddd/presentation/controllers"
)

func initHandlers(c commandBus.CommandBus, conn *pgxpool.Pool) {
	personRepository := repositoriesImpl.NewPersonRepository(conn)
	createPersonHandler := handlers.NewCreatePersonCommandHandler(personRepository)
	getPersonHandler := handlers.NewGetPersonCommandHandler(personRepository)
	c.RegisterHandler(reflect.TypeOf(person.GetPersonCommand{}), getPersonHandler)
	c.RegisterHandler(reflect.TypeOf(person.CreatePersonCommand{}), createPersonHandler)
}

func initRouter(c commandBus.CommandBus, a *fiber.App) {
	personController := controllers.NewPersonController("a", c)
	personRouter := routes.NewPersonRoutes(personController)
	personRouter.Init(a)
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

//func unMarshTest() {
//	var unmarshelPerson person.Person
//	email, err := personVO.NewEmail("test@gmail.com")
//	if err != nil {
//		log.Fatal(err)
//	}
//	name, err := personVO.NewName("firstName", "lastName")
//	if err != nil {
//		log.Fatal(err)
//	}
//	password, err := personVO.NewPassword("Pwd1234567u")
//	if err != nil {
//		log.Fatal(err)
//	}
//	person := person.NewPerson(name, email, password)
//	marshelPerson, err := json.Marshal(person)
//	if err != nil {
//		log.Fatal(err)
//	}
//	err = json.Unmarshal(marshelPerson, &unmarshelPerson)
//	fmt.Println(err)
//	fmt.Println(unmarshelPerson, "unmarshel Person")
//}

func main() {
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	}
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Live"})
	})
	commandBus := commandBus.NewCommandBus()
	initHandlers(commandBus, db.Conn())
	initRouter(commandBus, app)
	log.Fatal(app.Listen(":8080"))
}
