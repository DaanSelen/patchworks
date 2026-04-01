//go:build linux
// +build linux

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

func RunMeshbook(binPath, bookPath, targGroup string) (bool, string) {
	var args []string
	if len(bookPath) == 0 {
		args = []string{"--help"}
	} else {
		args = []string{"--nograce", "--indent", "-mb", bookPath, "--group", targGroup}
	}
	log.Printf("running with parameters: %v", args)

	cmd := exec.Command(binPath, args...)

	outputData, err := cmd.CombinedOutput()
	cleanData := ansi.ReplaceAllString(string(outputData), "")

	if err != nil {
		log.Printf("something went wrong when running the command: %v", err)
		log.Printf("captured output: %s", cleanData)

		return false, cleanData
	} else {
		return true, cleanData
	}
}
