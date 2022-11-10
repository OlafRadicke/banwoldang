package mediainformation

import (
	"os"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// CreateManifestLink Create a manifest file link in the link directory.
func (mediaInfo *MediaInformation) CreateManifestLink() {

	if mediaInfo.progConfig.UseHardLink == true {
		err := os.Link(mediaInfo.AbsoluteManifestSourcePath, mediaInfo.AbsoluteManifestLinkDirPath)
		if err != nil {
			cl.ErrorLogger.Println("Can't create link: ", err)
		}
	} else {
		err := os.Symlink(mediaInfo.AbsoluteManifestSourcePath, mediaInfo.AbsoluteManifestLinkDirPath)
		if err != nil {
			cl.ErrorLogger.Println("Can't create symlink: ", err)
		}
	}

}
