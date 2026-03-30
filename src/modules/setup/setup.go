package setup

import (
	"log"
	"os"
)

func PrepareEnvironment() (bool, error) {
	ok, err := fsCheck()

	return ok, err
}

func fsCheck() (bool, error) {
	dirs := []string{"./bin", "./books", "./history"}
	files := []string{"./api.conf", "./os_categories.json"}

	for _, d := range dirs {
		if _, err := os.Stat(d); err != nil {
			log.Printf("Unable to stat dir: %s", d)
			return false, err
		}
	}

	for _, f := range files {
		if _, err := os.Stat(f); err != nil {
			log.Printf("Unable to stat file: %s", f)
			return false, err
		}
	}

	return true, nil
}
