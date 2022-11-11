package mediainformation

import (
	"os"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// GenerateLinkDirTreeWithoutManifests Generate directory with symlinks of file without categories.
func (mediaInfo *MediaInformation) GenerateLinkDirTreeWithoutManifests() {

	mediaInfo.SetAbsoluteContentLinkDirPath("00-no-manifest")
	mediaInfo.SetAbsoluteManifestLinkDirPath("00-no-manifest")

	if _, err := os.Stat(mediaInfo.AbsoluteManifestSourcePath); !os.IsNotExist(err) {
		// path/to/whatever does not exist
		cl.InfoLogger.Println("Manifest file is exist. Skip media file handling. ", mediaInfo.AbsoluteManifestSourcePath)
		return
	}
	cl.InfoLogger.Println("Manifest file is NOT exist. Skip media file handling.", mediaInfo.AbsoluteManifestSourcePath)

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

	// mediaInfo.CreateNewEmptyManifestFile()

	err = os.Symlink(mediaInfo.AbsoluteManifestSourcePath, mediaInfo.AbsoluteManifestLinkDirPath)
	if err != nil {
		// cl.ErrorLogger.Fatal("Create symlink: ", err)
		cl.ErrorLogger.Println("Can't create symlink: ", err)
	}

}
