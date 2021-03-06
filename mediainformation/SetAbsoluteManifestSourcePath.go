package mediainformation

import (
	"path/filepath"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// SetAbsoluteManifestSourcePath Set the absolute source path of the Manifest file.
// @relativePath param is maybe the the relative path.
func (mediaInfo *MediaInformation) SetAbsoluteManifestSourcePath(relativePath string) {

	absoluteManifestSourcePath, err2 := filepath.Abs(relativePath)
	if err2 != nil {
		cl.ErrorLogger.Fatal(err2)
	}
	mediaInfo.AbsoluteManifestSourcePath = absoluteManifestSourcePath

}
