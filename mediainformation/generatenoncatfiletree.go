package mediainformation

import (
	"log"
	"os"
)

// Generate directory with symlinks of file without categories.
func (mediaInfo *MediaInformation) GenerateNonCatFileTree() {

	mediaInfo.SetAbsoluteContentLinkDirPath("00-no-cats")
	mediaInfo.SetAbsoluteManifestLinkDirPath("00-no-cats")

	katPath := mediaInfo.AbsoluteLinkDirPath + "/00-no-cats/" + "/.comments"
	if _, err := os.Stat(katPath); os.IsNotExist(err) {
		err := os.MkdirAll(katPath, 0770)
		if err != nil {
			log.Fatal(err)
		}
	}

	err := os.Symlink(mediaInfo.AbsoluteContentSourcePath, mediaInfo.AbsoluteContentLinkDirPath)
	if err != nil {
		// log.Fatal("Create symlink: ", err)
		log.Println("Create symlink: ", err)
	}

	log.Println("Chek: ", mediaInfo.AbsoluteManifestSourcePath)
	_, err = os.Stat(mediaInfo.AbsoluteManifestSourcePath)
	if os.IsNotExist(err) {
		log.Println("+++++++++++++++++++ CREAT MANIFEST +++++++++++++++++++++")
		log.Println(err, mediaInfo.AbsoluteManifestSourcePath)
		mediaInfo.CreateEmptyManifestFile()
	}

	err = os.Symlink(mediaInfo.AbsoluteManifestSourcePath, mediaInfo.AbsoluteManifestLinkDirPath)
	if err != nil {
		// log.Fatal("Create symlink: ", err)
		log.Println("Create symlink: ", err)
	}

}
