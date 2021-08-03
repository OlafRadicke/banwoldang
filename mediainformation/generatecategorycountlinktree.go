package mediainformation

import (
	"log"
	"strconv"
)

// Generate directory tree with symlinks of file depend of categories counts.
func (mediaInfo *MediaInformation) GenerateCategoryCountLinkTree() {

	log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	log.Println("++++++++++++++++++++++ has only ", len(mediaInfo.Categories), " category! ++++++++++++++++++++++")
	log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	catCountAsString := strconv.Itoa(len(mediaInfo.Categories))
	catSubDirectoryName := "00-cat-count/" + catCountAsString
	mediaInfo.SetAbsoluteContentLinkDirPath(catSubDirectoryName)
	mediaInfo.SetAbsoluteManifestLinkDirPath(catSubDirectoryName)
	mediaInfo.CreateLinkDirSubDir(catSubDirectoryName)
	mediaInfo.CreateContentLink()
	mediaInfo.CreateNewEmptyManifestFile()
	mediaInfo.CreateManifestLink()

}
