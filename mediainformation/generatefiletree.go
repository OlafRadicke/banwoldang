package mediainformation

import (
	"log"
	"os"
)

// Generate directory tree with symlinks of file with categories.
func (mediaInfo *MediaInformation) GenerateFileTree() {

	for i := 0; i < len(mediaInfo.Categories); i++ {
		mediaInfo.SetAbsoluteContentLinkDirPath(mediaInfo.Categories[i])
		mediaInfo.SetAbsoluteManifestLinkDirPath(mediaInfo.Categories[i])

		log.Println("mediaInfo.AbsoluteLinkDirPath: ", mediaInfo.AbsoluteLinkDirPath)
		katPath := mediaInfo.AbsoluteLinkDirPath + "/" + mediaInfo.Categories[i] + "/.comments"
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

		if _, err := os.Stat(mediaInfo.AbsoluteManifestSourcePath); os.IsNotExist(err) {
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
}
