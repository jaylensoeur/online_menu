package infrastructure

import (
	"encoding/json"
	"io"
	"menu/domain"
	"menu/domain/usecase"
	"os"
)

func readJsonFile[T any](path string) []T {
	currentDirectory, _ := os.Getwd()
	jsonFile, _ := os.Open(currentDirectory + path)

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {

		}
	}(jsonFile)

	byteValue, _ := io.ReadAll(jsonFile)
	var data []T

	err := json.Unmarshal(byteValue, &data)
	if err != nil {
		return nil
	}

	return data
}

func LoadJsonData(path string) map[string]*domain.Menu {

	jsonMenuDtos := readJsonFile[usecase.MenuDto](path)
	var data = map[string]*domain.Menu{}

	for _, menuDto := range jsonMenuDtos {
		data[menuDto.Uuid] = domain.NewMenu(
			domain.NewCafeId(
				domain.NewUuidWithUuid(menuDto.Uuid),
			),
			domain.NewTitle(menuDto.Title),
		)
	}

	return data
}
