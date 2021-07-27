package mediainformation

import (
	"log"
	"os"
)

// Generate directory tree with symlinks of file with categories.
func (mediaInfo *MediaInformation) GenerateFileTree() {

	for i := 0; i < len(mediaInfo.Categories); i++ {
		if len(mediaInfo.Categories) == 1 {
			log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
			log.Println("++++++++++++++++++++++++ has only one category! ++++++++++++++++++++++")
			log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

			mediaInfo.SetAbsoluteContentLinkDirPath("00-single-cats")
			mediaInfo.SetAbsoluteManifestLinkDirPath("00-single-cats")
			mediaInfo.CreateLinkDirSubDir("00-single-cats")

			err := os.Symlink(mediaInfo.AbsoluteContentSourcePath, mediaInfo.AbsoluteContentLinkDirPath)
			if err != nil {
				// log.Fatal("Create symlink: ", err)
				log.Println("Create symlink: ", err)
			}

			mediaInfo.CreateNewEmptyManifestFile()

			err = os.Symlink(mediaInfo.AbsoluteManifestSourcePath, mediaInfo.AbsoluteManifestLinkDirPath)
			if err != nil {
				// log.Fatal("Create symlink: ", err)
				log.Println("Create symlink: ", err)
			}

		}
		mediaInfo.SetAbsoluteContentLinkDirPath(mediaInfo.Categories[i])
		mediaInfo.SetAbsoluteManifestLinkDirPath(mediaInfo.Categories[i])
		mediaInfo.CreateLinkDirSubDir(mediaInfo.Categories[i])

		err := os.Symlink(mediaInfo.AbsoluteContentSourcePath, mediaInfo.AbsoluteContentLinkDirPath)
		if err != nil {
			// log.Fatal("Create symlink: ", err)
			log.Println("Create symlink: ", err)
		}

		mediaInfo.CreateNewEmptyManifestFile()

		err = os.Symlink(mediaInfo.AbsoluteManifestSourcePath, mediaInfo.AbsoluteManifestLinkDirPath)
		if err != nil {
			// log.Fatal("Create symlink: ", err)
			log.Println("Create symlink: ", err)
		}
	}
}
