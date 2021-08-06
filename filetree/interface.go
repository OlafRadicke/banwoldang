package filetree

import (
	"os"
)

type filetree interface {
	GoThroughCollection()
	PrintAll()
	fileHandler(path string, info os.FileInfo, err error) error

	SetAbsoluteSourcePath(string)
	SetAbsoluteLinkDir(string)
}

type FileTree struct {
	// The start path for searching media files
	SourcePath string
	// Location for the link directory
	LinkDir string
	// The number of founded manifest Files
	Findings int
	// Is the value true, than it will be create checksums as file names (for the links)
	UseChecksum bool
}
