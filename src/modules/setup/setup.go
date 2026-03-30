package setup

import "os"

func PrepareEnvironment() (bool, error) {
	if _, err := os.Stat("./bin"); err != nil {
		return false, err
	}

	if _, err := os.Stat("./books"); err != nil {
		return false, err
	}

	return true, nil
}
