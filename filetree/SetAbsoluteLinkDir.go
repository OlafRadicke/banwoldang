package filetree

import (
	"path/filepath"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// SetAbsoluteLinkDir set the absolute path of the target link directory
func (fileTree *FileTree) SetAbsoluteLinkDir(path string) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		cl.ErrorLogger.Fatal(err)
	}
	fileTree.LinkDir = absolutePath
}
