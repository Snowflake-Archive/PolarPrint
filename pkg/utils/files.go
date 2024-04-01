package utils

import (
	"log"
	"os"
)

func GetFilenames() []string {
	files, err := os.ReadDir("./prints")
	if err != nil {
		log.Fatal(err)
	}
	fileNames := []string{}
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	return fileNames
}
