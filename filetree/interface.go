package filetree

import (
	"os"
)

// filetree The inteface of the feletree struckt
type filetree interface {
	GoThroughCollection()
	NewFileTree() FileTree
	PrintAll()

	fileHandler(path string, info os.FileInfo, err error) error

	SetAbsoluteSourcePath(string)
	SetAbsoluteLinkDir(string)
}

// FileTree This struct represent the information about the tree
type FileTree struct {
	// SourcePath The start path for searching media files
	SourcePath string

	// LinkDir Location for the link directory
	LinkDir string

	// Findings The number of founded manifest Files
	Findings int

	// UseChecksum Is the value true, than it will be create checksums as file names (for the links)
	UseChecksum bool

	// AllUsedCategories A map with all used categories.
	AllUsedCategories map[string]int
}

// NewProgArguments create new instance of FileTree and get it back.
func NewFileTree() FileTree {
	fileTree := FileTree{}

	fileTree.SourcePath = ""
	fileTree.LinkDir = ""
	fileTree.Findings = 0
	fileTree.UseChecksum = false
	fileTree.AllUsedCategories = make(map[string]int)
	return fileTree
}
