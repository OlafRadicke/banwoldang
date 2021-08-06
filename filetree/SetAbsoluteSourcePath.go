package filetree

import (
	"log"
	"path/filepath"
)

// SetAbsoluteSourcePath set the absolute start path from the source directory
func (fileTree *FileTree) SetAbsoluteSourcePath(path string) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}
	fileTree.SourcePath = absolutePath
}
