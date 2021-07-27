package mediainformation

import (
	"log"
	"os"
)

// Create a sub directory in the link tree directory. Include the
// sub directory for the manifests (.comments)
// @subDir sing with the mane of the new sub directory.
func (mediaInfo *MediaInformation) CreateLinkDirSubDir(subDir string) {

	log.Println("mediaInfo.AbsoluteLinkDirPath: ", mediaInfo.AbsoluteLinkDirPath)
	katPath := mediaInfo.AbsoluteLinkDirPath + "/" + subDir + "/.comments"
	_, err := os.Stat(katPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(katPath, 0770)
		if err != nil {
			log.Fatal(err)
		}
	}

}
