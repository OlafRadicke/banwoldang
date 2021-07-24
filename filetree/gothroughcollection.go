package filetree

import (
	"log"
	"path/filepath"
)

func (fileTree *FileTree) GoThroughCollection() {
	err := filepath.Walk(fileTree.StartPath, fileTree.fileHandler)
	if err != nil {
		log.Println(err)
	}
}
