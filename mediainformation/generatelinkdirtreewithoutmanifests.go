package mediainformation

import (
	"os"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// Generate directory with symlinks of file without categories.
func (mediaInfo *MediaInformation) GenerateLinkDirTreeWithoutManifests() {

	mediaInfo.SetAbsoluteContentLinkDirPath("00-no-manifest")
	mediaInfo.SetAbsoluteManifestLinkDirPath("00-no-manifest")

	katPath := mediaInfo.AbsoluteLinkDirPath + "/00-no-manifest/" + "/.comments"
	if _, err := os.Stat(katPath); os.IsNotExist(err) {
		err := os.MkdirAll(katPath, 0770)
		if err != nil {
			cl.ErrorLogger.Fatal(err)
		}
	}

	err := os.Symlink(mediaInfo.AbsoluteContentSourcePath, mediaInfo.AbsoluteContentLinkDirPath)
	if err != nil {
		// cl.ErrorLogger.Fatal("Create symlink: ", err)
		cl.ErrorLogger.Println("Can't create symlink: ", err)
	}

	mediaInfo.CreateNewEmptyManifestFile()

	err = os.Symlink(mediaInfo.AbsoluteManifestSourcePath, mediaInfo.AbsoluteManifestLinkDirPath)
	if err != nil {
		// cl.ErrorLogger.Fatal("Create symlink: ", err)
		cl.ErrorLogger.Println("Can't create symlink: ", err)
	}

}
