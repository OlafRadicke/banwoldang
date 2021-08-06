package mediainformation

import (
	"os"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// Create a sub directory in the link tree directory. Include the
// sub directory for the manifests (.comments)
// @subDir sing with the mane of the new sub directory.
func (mediaInfo *MediaInformation) CreateLinkDirSubDir(subDir string) {
	katPath := mediaInfo.AbsoluteLinkDirPath + "/" + subDir + "/.comments"
	_, err := os.Stat(katPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(katPath, 0770)
		if err != nil {
			cl.ErrorLogger.Fatal(err)
		}
	}

}
