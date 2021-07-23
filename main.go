package main

import (
	"log"
	"os"
	"path/filepath"

	common "github.com/OlafRadicke/banwoldang/common"
)

// Struct with command line arguments
type ProgArguments struct {
	// Location with the media files
	MediaDir string
	// Location for the generic directory
	GenericDir string
	// The minimum of needed params
	MinimumArguments int
}

func main() {
	progArguments := checkInput()

	fileTree := common.FileTree{}
	absolutStartPath, err := filepath.Abs(progArguments.MediaDir)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("absolutStartPath: ", absolutStartPath)
	fileTree.StartPath = absolutStartPath
	absoluteGenericDir, err := filepath.Abs(progArguments.GenericDir)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("absoluteGenericDir: ", absoluteGenericDir)
	fileTree.GenericDir = absoluteGenericDir
	log.Println("Search in: ", fileTree.StartPath)

	fileTree.GoThroughCollection()
	log.Println("Findings: ", fileTree.Findings)
}

func checkInput() *ProgArguments {
	progArguments := ProgArguments{}
	givenArguments := len(os.Args)
	progArguments.MinimumArguments = 3
	if givenArguments < progArguments.MinimumArguments {
		log.Fatal("To less arguments! ", progArguments.MinimumArguments, " needed and ", givenArguments, "given ")
	} else {
		// programName := os.Args[0]
		progArguments.MediaDir = os.Args[1]
		progArguments.GenericDir = os.Args[2]

		// startPasth = os.Args[1]
		// log.Println("Arg numbers: ", givenArguments)
		// log.Println("Starting the application: ", programName)
	}

	return &progArguments
}
