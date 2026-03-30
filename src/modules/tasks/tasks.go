package tasks

import (
	"log"
	"os"
	"strings"
)

func isYaml(filename string) bool {
	return strings.HasSuffix(filename, ".yaml") || strings.HasSuffix(filename, ".yml")
}

func ListAvailableBooks() ([]string, error) {
	files, err := os.ReadDir("./books")
	if err != nil {
		log.Printf("failed to read the './books' directory: %v", err)
		return []string{}, err
	}

	foundBooks := []string{}
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		fName := f.Name()
		log.Println(fName)
		if isYaml(fName) {
			fullRelPath := "./books/" + fName
			foundBooks = append(foundBooks, fullRelPath)
		}
	}

	return foundBooks, nil
}
