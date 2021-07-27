package main

import (
	"log"
	"os"
	"path/filepath"

	filetree "github.com/OlafRadicke/banwoldang/filetree"
)

// Struct with command line arguments
type ProgArguments struct {
	// Location with the media files
	MediaDir string
	// Location for the link directory
	LinkDir string
	// The minimum of needed params
	MinimumArguments int
}

func main() {
	progArguments := checkInput()

	fileTree := filetree.FileTree{}
	absolutStartPath, err := filepath.Abs(progArguments.MediaDir)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("absolutStartPath: ", absolutStartPath)
	fileTree.StartPath = absolutStartPath
	absoluteLinkDir, err := filepath.Abs(progArguments.LinkDir)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("absoluteLinkDir: ", absoluteLinkDir)
	fileTree.LinkDir = absoluteLinkDir
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
		progArguments.MediaDir = os.Args[1]
		progArguments.LinkDir = os.Args[2]
	}

	return &progArguments
}
