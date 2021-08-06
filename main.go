package main

import (
	"log"
	"os"
	"strings"

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
}

// NewProgArguments create new instance of ProgArguments
func NewProgArguments() ProgArguments {
	progArguments := ProgArguments{}
	progArguments.SourceDir = ""
	progArguments.LinkDir = ""
	progArguments.MinimumArguments = 3
	progArguments.UseChecksum = false
	return progArguments
}

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	infoFile, err := os.OpenFile("info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	errorFile, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(infoFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {

	// InfoLogger.Println("der neue Info-Logger")
	// ErrorLogger.Println("der neue Error-Logger")
	// ErrorLogger.Fatal("der neue Error-Logger mit hartem Ende...")

	progArguments := checkInput()

	fileTree := filetree.FileTree{}
	fileTree.SetAbsoluteSourcePath(progArguments.SourceDir)
	fileTree.SetAbsoluteLinkDir(progArguments.LinkDir)
	log.Println("absoluteLinkDir: ", fileTree.LinkDir)
	log.Println("Search in: ", fileTree.SourcePath)

	fileTree.GoThroughCollection()
	log.Println("Count of founded files: ", fileTree.Findings)
}

// checkInput read the programme arguments
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
			progArguments.SourceDir = argParts[1]
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
