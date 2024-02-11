package restapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"menu/domain/create"
	"menu/domain/list"
	"menu/domain/single"
	"menu/infrastructure/repository/mongo"
	"os"
)

func Run() {

	err := godotenv.Load(".env")
	if err != nil {
		return
	}
	mongoDao := mongo.NewMongoDao(os.Getenv("MONGODB_URI"))
	mongoDao.Connect()

	defer mongoDao.Disconnect()
	repository := mongo.NewMenuMongoRepository(mongoDao)

	controller := NewMenuController(
		create.NewCreateMenu(repository),
		single.NewGetMenu(repository),
		list.NewListMenu(repository),
	)
	app := fiber.New()
	app.Post("menus", controller.Add())
	app.Get("menus", controller.List())
	app.Get("menus/:id", controller.Get())

	err = app.Listen(":8080")

	if err != nil {
		return
	}
}
