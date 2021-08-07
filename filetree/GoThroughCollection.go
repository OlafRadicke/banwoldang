package filetree

import (
	"path/filepath"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// GoThroughCollection This function gos through the source directory recursively.
func (fileTree *FileTree) GoThroughCollection() {
	err := filepath.Walk(fileTree.SourcePath, fileTree.fileHandler)
	if err != nil {
		cl.ErrorLogger.Println(err)
	}
}
