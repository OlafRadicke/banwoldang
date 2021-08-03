package mediainformation

import (
	"log"
)

// Generate directory tree with symlinks of file with categories.
func (mediaInfo *MediaInformation) GenerateLinkDirTree() {

	mediaInfo.GenerateOldNamePartLinkTree()

	if len(mediaInfo.Categories) > 1 {
		log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		log.Println("+++++++++++++++++++ this file has ", len(mediaInfo.Categories), " categories ++++++++++++++++++++++")
		log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	}

	for i := 0; i < len(mediaInfo.Categories); i++ {
		if len(mediaInfo.Categories) < 5 {

			mediaInfo.GenerateCategoryCountLinkTree()
		}

		mediaInfo.SetAbsoluteContentLinkDirPath(mediaInfo.Categories[i])
		mediaInfo.SetAbsoluteManifestLinkDirPath(mediaInfo.Categories[i])
		mediaInfo.CreateLinkDirSubDir(mediaInfo.Categories[i])
		mediaInfo.CreateContentLink()
		mediaInfo.CreateNewEmptyManifestFile()
		mediaInfo.CreateManifestLink()

	}

}
