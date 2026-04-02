package runner

import (
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
)

var ansi = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func FindMeshbookBinary() (bool, string) {
	var osBin string

	switch runtime.GOOS {
	case "windows":
		osBin = "meshbook.exe"
	case "linux":
		osBin = "meshbook"
	default:
		log.Println("undefined operating system")
	}

	log.Println("going to search for:", osBin)

	binaryFound := false
	var binaryPath string
	for _, f := range []string{("./" + osBin), ("./bin/" + osBin)} {
		objInfo, err := os.Stat(f)

		if err == nil && objInfo.Mode().IsRegular() {
			binaryFound = true
			binaryPath = f
			log.Printf("found binary at %s", f)
			break
		}
	}

	if binaryFound {
		return true, binaryPath
	} else {
		log.Println("binary not found!")
		return false, ""
	}
}

func RunMeshbook(binPath, bookPath string, silent bool, targGroup string) (bool, string) {
	// Meshbook argument compilation
	var args []string
	if len(bookPath) == 0 {
		args = []string{"--help"}
	} else {
		args = []string{"--nograce", "--indent", "-mb", bookPath, "--group", targGroup}
	}

	if silent {
		args = append(args, "--silent")
	}

	// Display what we are about to be running
	log.Printf("running with parameters: %v", args)

	// Actually spawn the process
	cmd := exec.Command(binPath, args...)

	// Capture stdout & stderr
	outputData, err := cmd.CombinedOutput()
	cleanData := ansi.ReplaceAllString(string(outputData), "")

	log.Println("evaluating returned state")
	if err != nil {
		log.Printf("something went wrong when running the command: %v", err)
		log.Printf("captured output: %s", cleanData)

		return false, cleanData
	} else {
		log.Printf("captured output: %s", cleanData)

		return true, cleanData
	}
}
