package utils

import (
	"os"
)

func GetFilenames() ([]string, error) {
	files, err := os.ReadDir("./prints")

	if err != nil {
		return nil, err
	}

	fileNames := []string{}

	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	return fileNames, nil
}
