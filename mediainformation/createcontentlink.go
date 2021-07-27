package mediainformation

import (
	"log"
	"os"
)

// Create a content file link in the link directory.
func (mediaInfo *MediaInformation) CreateContentLink() {

	err := os.Symlink(mediaInfo.AbsoluteContentSourcePath, mediaInfo.AbsoluteContentLinkDirPath)
	if err != nil {
		// log.Fatal("Create symlink: ", err)
		log.Println("Create symlink: ", err)
	}

}
