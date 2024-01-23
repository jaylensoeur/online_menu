package restapi

import (
	"fmt"
	"menu/application/create_menu"
	"menu/infrastructure/mongo"
	"menu/infrastructure/restapi/menu"
	"net/http"
)

func Run(port string) {

	menuController := menu.NewMenuController(
		create_menu.NewCreateMenu(
			mongo.NewMenuMongoRepository(),
		),
	)

	http.HandleFunc("/add", menuController.Add([]string{"post"}))

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("error %v", err)
		return
	}
}
