package filetree

import (
	"path/filepath"
    "fmt"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// SetAbsoluteLinkDir set the absolute path of the target link directory
func (fileTree *FileTree) SetAbsoluteLinkDir(path string) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		fmt.Println("error! check the log files for more information")
		cl.ErrorLogger.Fatal(err)
	}
	fileTree.LinkDir = absolutePath
}
