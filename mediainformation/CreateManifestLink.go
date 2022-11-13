package mediainformation

import (
	"os"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// CreateManifestLink Create a manifest file link in the link directory.
func (mediaInfo *MediaInformation) CreateManifestLink() {

	if mediaInfo.progConfig.UseHardLink == true {
		err := os.Link(mediaInfo.Comments.FilePath, mediaInfo.AbsoluteManifestLinkDirPath)
		if err != nil {
			cl.ErrorLogger.Println("Can't create link: ", err)
		}
	} else {
		cl.ErrorLogger.Println("Try to link: ")
		cl.ErrorLogger.Println(mediaInfo.Comments.FilePath)
		cl.ErrorLogger.Println(mediaInfo.AbsoluteManifestLinkDirPath)
		err := os.Symlink(mediaInfo.Comments.FilePath, mediaInfo.AbsoluteManifestLinkDirPath)
		if err != nil {
			cl.ErrorLogger.Println("Can't create symlink: ", err)
		}
	}

}
