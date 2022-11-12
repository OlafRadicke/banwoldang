package linkdirectories

import (
	"strconv"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
	"github.com/OlafRadicke/banwoldang/mediainformation"
)

// GenerateLinkDirTreeOfCategoryCount Generate directory tree with symlinks of file
// depend of categories counts.
func GenerateLinkDirTreeOfCategoryCount(mediaInfo *mediainformation.MediaInformation) {

	cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	cl.InfoLogger.Println("++++++++++++++++++++++ has only ", len(mediaInfo.Comments.GetCategories()), " category! ++++++++++++++++++++++")
	cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	catCountAsString := strconv.Itoa(len(mediaInfo.Comments.GetCategories()))
	catSubDirectoryName := "00-cat-count/" + catCountAsString
	mediaInfo.SetAbsoluteContentLinkDirPath(catSubDirectoryName)
	mediaInfo.SetAbsoluteManifestLinkDirPath(catSubDirectoryName)
	mediaInfo.CreateLinkDirSubDir(catSubDirectoryName)
	mediaInfo.CreateContentLink()
	// mediaInfo.CreateNewEmptyManifestFile()
	mediaInfo.CreateManifestLink()

}
