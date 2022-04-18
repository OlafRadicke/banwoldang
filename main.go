package main

import (
	"os"
	"strings"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
	filetree "github.com/OlafRadicke/banwoldang/filetree"
)

// ProgArguments struct with command line arguments
type ProgArguments struct {
	// Location with the media files
	SourceDir string
	// Location for the link directory
	LinkDir string
	// The minimum of needed params
	MinimumArguments int
	// Is the value true, than it will be create checksums as file names (for the links)
	UseChecksum bool
	// Is the value true, than it will be try to create hard links.
	UseHardLink bool
}

// NewProgArguments create new instance of ProgArguments
func NewProgArguments() ProgArguments {
	progArguments := ProgArguments{}
	progArguments.SourceDir = ""
	progArguments.LinkDir = ""
	progArguments.MinimumArguments = 3
	progArguments.UseChecksum = false
	progArguments.UseHardLink = false
	return progArguments
}

func main() {

	cl.InfoLogger.Println("================================= PROGRAMM START ==============================")
	cl.ErrorLogger.Println("================================ PROGRAMM START ==============================")

	// cl.InfoLogger.Println("PROGRAMM START")
	// cl.ErrorLogger.Println("PROGRAMM START", 1, "test")
	// cl.ErrorLogger.Fatal("der neue Error-Logger mit hartem Ende...")

	progArguments := checkInput()

	// fileTree := filetree.FileTree{}
	fileTree := filetree.NewFileTree()
	fileTree.SetAbsoluteSourcePath(progArguments.SourceDir)
	fileTree.SetAbsoluteLinkDir(progArguments.LinkDir)
	fileTree.UseChecksum = progArguments.UseChecksum
	fileTree.UseHardLink = progArguments.UseHardLink
	cl.InfoLogger.Println("absoluteLinkDir: ", fileTree.LinkDir)
	cl.InfoLogger.Println("Search in: ", fileTree.SourcePath)

	fileTree.GoThroughCollection()
	cl.InfoLogger.Println("Count of founded files: ", fileTree.Findings)
	cl.InfoLogger.Println("Count of founded categories: ", len(fileTree.AllUsedCategories))
	fileTree.CreateTagsXmlFile()
}

// checkInput read the programme arguments
func checkInput() *ProgArguments {
	progArguments := NewProgArguments()
	givenArguments := len(os.Args)
	if givenArguments < progArguments.MinimumArguments {
		cl.ErrorLogger.Fatal("To less arguments! ", progArguments.MinimumArguments, " needed and ", givenArguments, "given ")
	}

	for i := 1; i < len(os.Args); i++ {
		argParts := strings.Split(os.Args[i], "=")
		if len(argParts) != 2 {
			cl.ErrorLogger.Fatal("Argument is wrong: ", os.Args[i])
		}
		cl.InfoLogger.Println("name: ", argParts[0])
		cl.InfoLogger.Println("volume: ", argParts[1])

		switch argParts[0] {
		case "--source-dir":
			progArguments.SourceDir = argParts[1]
		case "--link-dir":
			progArguments.LinkDir = argParts[1]
		case "--checksum":
			if argParts[1] == "true" {
				progArguments.UseChecksum = true
			} else {
				progArguments.UseChecksum = false
			}
		case "--hardlinks":
			if argParts[1] == "true" {
				progArguments.UseHardLink = true
			} else {
				progArguments.UseHardLink = false
			}
		default:
			cl.ErrorLogger.Fatal("Unknown parameter: ", argParts[0])
		}

		cl.InfoLogger.Println("UseHardLink is: ", progArguments.UseHardLink)
		cl.InfoLogger.Println("UseChecksum is: ", progArguments.UseChecksum)
	}
	return &progArguments
}
