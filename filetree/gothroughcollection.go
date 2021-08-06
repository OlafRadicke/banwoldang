package filetree

import (
	"path/filepath"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

func (fileTree *FileTree) GoThroughCollection() {
	err := filepath.Walk(fileTree.SourcePath, fileTree.fileHandler)
	if err != nil {
		cl.InfoLogger.Println(err)
	}
}
