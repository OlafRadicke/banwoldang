package filetree

import (
	"log"
	"path/filepath"
)

// SetAbsoluteStartPath set the absolute start path from the source directory
func (fileTree *FileTree) SetAbsoluteStartPath(path string) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}
	fileTree.StartPath = absolutePath
}
