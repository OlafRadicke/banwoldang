package mediainformation

import (
	"path/filepath"
    "fmt"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// SetAbsoluteManifestSourcePath Set the absolute source path of the Manifest file.
// @relativePath param is maybe the the relative path.
func (mediaInfo *MediaInformation) SetAbsoluteManifestSourcePath(relativePath string) {

	absoluteManifestSourcePath, err2 := filepath.Abs(relativePath)
	if err2 != nil {
		cl.ErrorLogger.Fatal(err2)
	}
	if filepath.Ext(absoluteManifestSourcePath) != ".xml" {
		fmt.Println("error! check the log files for more information")
		cl.ErrorLogger.Fatal("Thats a not a right name for an command file: ", absoluteManifestSourcePath)
	}
	mediaInfo.Comments.FilePath = absoluteManifestSourcePath

}
