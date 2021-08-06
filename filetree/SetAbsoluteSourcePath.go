package filetree

import (
	"path/filepath"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// SetAbsoluteSourcePath set the absolute start path from the source directory
func (fileTree *FileTree) SetAbsoluteSourcePath(path string) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		cl.ErrorLogger.Fatal(err)
	}
	fileTree.SourcePath = absolutePath
}
