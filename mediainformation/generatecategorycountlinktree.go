package mediainformation

import "log"

// Generate directory tree with symlinks of file depend of categories counts.
func (mediaInfo *MediaInformation) GenerateCategoryCountLinkTree(maxCount int) {

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
