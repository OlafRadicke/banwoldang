package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	filetree "github.com/OlafRadicke/banwoldang/filetree"
)

// ProgArguments struct with command line arguments
type ProgArguments struct {
	// Location with the media files
	MediaDir string
	// Location for the link directory
	LinkDir string
	// The minimum of needed params
	MinimumArguments int
	// Is the value true, than it will be create checksums as file names (for the links)
	UseChecksum bool
}

// NewProgArguments create new instance of ProgArguments
func NewProgArguments() ProgArguments {
	progArguments := ProgArguments{}
	progArguments.MediaDir = ""
	progArguments.LinkDir = ""
	progArguments.MinimumArguments = 3
	progArguments.UseChecksum = false
	return progArguments
}

func main() {
	progArguments := checkInput()

	fileTree := filetree.FileTree{}
	absolutStartPath, err := filepath.Abs(progArguments.MediaDir)
	if err != nil {
		log.Fatal(err)
	}
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

// ProgArguments read the programme arguments
func checkInput() *ProgArguments {
	progArguments := NewProgArguments()
	givenArguments := len(os.Args)
	if givenArguments < progArguments.MinimumArguments {
		log.Fatal("To less arguments! ", progArguments.MinimumArguments, " needed and ", givenArguments, "given ")
	}

	for i := 1; i < len(os.Args); i++ {
		argParts := strings.Split(os.Args[i], "=")
		if len(argParts) != 2 {
			log.Fatal("Argument is wrong: ", os.Args[i])
		}
		log.Println("name: ", argParts[0])
		log.Println("volume: ", argParts[1])

		switch argParts[0] {
		case "--source-dir":
			progArguments.MediaDir = argParts[1]
		case "--link-dir":
			progArguments.LinkDir = argParts[1]
		case "--checksum":
			if argParts[1] == "true" {
				progArguments.UseChecksum = true
			} else {
				progArguments.UseChecksum = false
			}
		default:
			log.Fatal("Unknown parameter: ", argParts[0])
		}

		log.Println("UseChecksum is: ", progArguments.UseChecksum)
	}
	return &progArguments
}
