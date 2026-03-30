package setup

import (
	"fmt"
	"log"
	"os"
)

func PrepareEnvironment() (bool, error) {
	if ok, err := fsCheck(); !ok || err != nil {
		return ok, err
	}

	if ok, err := ensState(); !ok || err != nil {
		return ok, err
	}

	log.Println("Validated or made state compliant")
	return true, nil
}

// Filesystem check
func fsCheck() (bool, error) {
	dirs := []string{"./bin", "./books"}

	for _, d := range dirs {
		if _, err := os.Stat(d); err != nil {
			log.Printf("unable to stat dir: %s, trying to make it...", d)

			err := os.Mkdir(d, 0755)
			if err != nil {
				return false, err
			}
		}
	}

	return true, nil
}

// Basic state
func ensState() (bool, error) {
	files := []string{"./api.conf", "./os_categories.json",
		"./books/rdpCheck.yaml",
		"./books/updateAptCache.yaml", "./books/updateOs.yaml",
		"./books/enableVncConsent.yaml", "./books/disableVncConsent.yaml",
	}

	failedState := false
	for _, f := range files {
		if _, err := os.Stat(f); err != nil {
			log.Printf("unable to stat: %s, creating it from template...", f)

			var err error
			switch f {
			case "./api.conf":
				err = os.WriteFile(f, []byte(apiConf), 0644)
			case "./os_categories.json":
				err = os.WriteFile(f, []byte(osCategories), 0644)
			case "./books/rdpCheck.yaml":
				err = os.WriteFile(f, []byte(rdpCheck), 0644)
			case "./books/updateAptCache.yaml":
				err = os.WriteFile(f, []byte(updateAptCache), 0644)
			case "./books/updateOs.yaml":
				err = os.WriteFile(f, []byte(updateOs), 0644)
			case "./books/enableVncConsent.yaml":
				err = os.WriteFile(f, []byte(enableVncConsent), 0644)
			case "./books/disableVncConsent.yaml":
				err = os.WriteFile(f, []byte(disableVncConsent), 0644)
			default:
				log.Println("no template defined for this file... not making it")
			}

			if err != nil {
				log.Printf("error trying to create from template: %v", err)
				failedState = true
			}
		}
	}

	if failedState {
		return false, fmt.Errorf("failed to ensure the basic operating state. see above for details")
	} else {
		return true, nil
	}
}
