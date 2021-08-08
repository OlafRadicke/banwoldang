package mediainformation

import (
	"strconv"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// GenerateLinkDirTreeOfCategoryCount Generate directory tree with symlinks of file
// depend of categories counts.
func (mediaInfo *MediaInformation) GenerateLinkDirTreeOfCategoryCount() {

	cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	cl.InfoLogger.Println("++++++++++++++++++++++ has only ", len(mediaInfo.Categories), " category! ++++++++++++++++++++++")
	cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	catCountAsString := strconv.Itoa(len(mediaInfo.Categories))
	catSubDirectoryName := "00-cat-count/" + catCountAsString
	mediaInfo.SetAbsoluteContentLinkDirPath(catSubDirectoryName)
	mediaInfo.SetAbsoluteManifestLinkDirPath(catSubDirectoryName)
	mediaInfo.CreateLinkDirSubDir(catSubDirectoryName)
	mediaInfo.CreateContentLink()
	mediaInfo.CreateNewEmptyManifestFile()
	mediaInfo.CreateManifestLink()

}
