package filetree

import (
	"log"
	"path/filepath"
)

// SetAbsoluteLinkDir set the absolute path of the target link directory
func (fileTree *FileTree) SetAbsoluteLinkDir(path string) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}
	fileTree.LinkDir = absolutePath
}
