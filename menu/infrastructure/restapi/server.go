package restapi

import (
	"github.com/gofiber/fiber/v2"
	"menu/domain/usecase"
	"menu/infrastructure"
	"menu/infrastructure/repository/inmemory"
	"menu/infrastructure/restapi/menu"
)

func Run() {

	inMemoryRepository := inmemory.NewInMemoryRepository(infrastructure.LoadJsonData("/infrastructure/repository/inmemory/MOCK_DATA.json"))

	controller := menu.NewMenuController(
		usecase.NewCreateMenu(inMemoryRepository),
		usecase.NewGetMenu(inMemoryRepository),
		usecase.NewListMenu(inMemoryRepository),
	)

	app := fiber.New()
	app.Post("menus", controller.Add())
	app.Get("menus", controller.List())
	app.Get("menus/:id", controller.Get())

	err := app.Listen(":8080")

	if err != nil {
		return
	}
}
