package mediainformation

import (
	"path/filepath"
    "fmt"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// SetAbsoluteContentSourcePath Set the absolute source path of the conten file.
// @relativePath param is maybe the the relative path.
func (mediaInfo *MediaInformation) SetAbsoluteContentSourcePath(relativePath string) {

	absoluteContentSourcePath, err2 := filepath.Abs(relativePath)
	if err2 != nil {
		fmt.Println("error! check the log files for more information")
		cl.ErrorLogger.Fatal(err2)
	}
	mediaInfo.AbsoluteContentSourcePath = absoluteContentSourcePath
}
