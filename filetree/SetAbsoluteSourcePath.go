package filetree

import (
	"path/filepath"
    "fmt"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// SetAbsoluteSourcePath set the absolute start path from the source directory
func (fileTree *FileTree) SetAbsoluteSourcePath(path string) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		fmt.Println("error! check the log files for more information")
		cl.ErrorLogger.Fatal(err)
	}
	fileTree.SourcePath = absolutePath
}
