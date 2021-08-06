package mediainformation

import (
	"path/filepath"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// Set the absolute source path of the conten file.
// @relativePath param is maybe the the relative path.
func (mediaInfo *MediaInformation) SetAbsoluteContentSourcePath(relativePath string) {

	absoluteContentSourcePath, err2 := filepath.Abs(relativePath)
	if err2 != nil {
		cl.ErrorLogger.Fatal(err2)
	}
	mediaInfo.AbsoluteContentSourcePath = absoluteContentSourcePath
}
