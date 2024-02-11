package infrastructure

import (
	"encoding/json"
	"io"
	"os"
)

func ReadJsonFile[T any](path string) []T {
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
