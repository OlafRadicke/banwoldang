package mediainformation

import (
	"os"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// Create a manifest file link in the link directory.
func (mediaInfo *MediaInformation) CreateManifestLink() {

	err := os.Symlink(mediaInfo.AbsoluteManifestSourcePath, mediaInfo.AbsoluteManifestLinkDirPath)
	if err != nil {
		cl.ErrorLogger.Println("Can't create symlink: ", err)
	}

}
