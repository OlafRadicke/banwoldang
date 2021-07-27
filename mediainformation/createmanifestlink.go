package mediainformation

import (
	"log"
	"os"
)

// Create a manifest file link in the link directory.
func (mediaInfo *MediaInformation) CreateManifestLink() {

	err := os.Symlink(mediaInfo.AbsoluteManifestSourcePath, mediaInfo.AbsoluteManifestLinkDirPath)
	if err != nil {
		// log.Fatal("Create symlink: ", err)
		log.Println("Create symlink: ", err)
	}

}
