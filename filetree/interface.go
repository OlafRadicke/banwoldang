package filetree

import (
	"os"
)

type filetree interface {
	GoThroughCollection()
	PrintAll()
	fileHandler(path string, info os.FileInfo, err error) error
}

type FileTree struct {
	// The start path for searching media files
	StartPath string
	// Location for the generic directory
	GenericDir string
	// The number of founded manifest Files
	Findings int
}
