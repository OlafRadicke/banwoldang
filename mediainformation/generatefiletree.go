package mediainformation

import (
	"log"
)

// Generate directory tree with symlinks of file with categories.
func (mediaInfo *MediaInformation) GenerateLinkDirTree() {

	listOfParts := mediaInfo.ExtractFileNameParts()
	for i := 0; i < len(listOfParts); i++ {

		log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		log.Println("++++++++++++++++++++++++ create old name category+++++++++++++++++++++")
		log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

		mediaInfo.SetAbsoluteContentLinkDirPath("00-olad-name-part/" + listOfParts[i])
		mediaInfo.SetAbsoluteManifestLinkDirPath("00-olad-name-part/" + listOfParts[i])
		mediaInfo.CreateLinkDirSubDir("00-olad-name-part/" + listOfParts[i])
		mediaInfo.CreateContentLink()
		mediaInfo.CreateNewEmptyManifestFile()
		mediaInfo.CreateManifestLink()
	}

	if len(mediaInfo.Categories) > 1 {
		log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		log.Println("+++++++++++++++++++ this file has ", len(mediaInfo.Categories), " categories ++++++++++++++++++++++")
		log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	}

	for i := 0; i < len(mediaInfo.Categories); i++ {
		if len(mediaInfo.Categories) == 1 {
			log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
			log.Println("++++++++++++++++++++++++ has only one category! ++++++++++++++++++++++")
			log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

			mediaInfo.SetAbsoluteContentLinkDirPath("00-single-cats")
			mediaInfo.SetAbsoluteManifestLinkDirPath("00-single-cats")
			mediaInfo.CreateLinkDirSubDir("00-single-cats")
			mediaInfo.CreateContentLink()
			mediaInfo.CreateNewEmptyManifestFile()
			mediaInfo.CreateManifestLink()

		}

		mediaInfo.SetAbsoluteContentLinkDirPath(mediaInfo.Categories[i])
		mediaInfo.SetAbsoluteManifestLinkDirPath(mediaInfo.Categories[i])
		mediaInfo.CreateLinkDirSubDir(mediaInfo.Categories[i])
		mediaInfo.CreateContentLink()
		mediaInfo.CreateNewEmptyManifestFile()
		mediaInfo.CreateManifestLink()

	}

}
