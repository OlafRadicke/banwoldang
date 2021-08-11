package mediainformation

import (
	"os"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// CreateContentLink Create a content file link in the link directory.
func (mediaInfo *MediaInformation) CreateContentLink() error {

	err := os.Symlink(mediaInfo.AbsoluteContentSourcePath, mediaInfo.AbsoluteContentLinkDirPath)
	if err != nil {
		// cl.ErrorLogger.Fatal("Create symlink: ", err)
		cl.ErrorLogger.Println("Can't create symlink: ", err)
		return err
	}
	return nil
}
