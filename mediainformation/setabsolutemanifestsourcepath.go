package mediainformation

import (
	"log"
	"path/filepath"
)

// Set the absolute source path of the Manifest file.
// @relativePath param is maybe the the relative path.
func (mediaInfo *MediaInformation) SetAbsoluteManifestSourcePath(relativePath string) {

	absoluteManifestSourcePath, err2 := filepath.Abs(relativePath)
	if err2 != nil {
		log.Fatal(err2)
	}
	mediaInfo.AbsoluteManifestSourcePath = absoluteManifestSourcePath

}
