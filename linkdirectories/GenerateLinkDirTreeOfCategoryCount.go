package linkdirectories

import (
	"strconv"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// GenerateLinkDirTreeOfCategoryCount Generate directory tree with symlinks of file
// depend of categories counts.
func (linkdirectories  *Linkdirectories) GenerateLinkDirTreeOfCategoryCount() {

	cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	cl.InfoLogger.Println("++++++++++++++++++++++ has only ", len(linkdirectories.getCategories()), " category! ++++++++++++++++++++++")
	cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	catCountAsString := strconv.Itoa(len(linkdirectories.getCategories()))
	catSubDirectoryName := "00-cat-count/" + catCountAsString
	linkdirectories.mediaInfo.SetAbsoluteContentLinkDirPath(catSubDirectoryName)
	linkdirectories.mediaInfo.SetAbsoluteManifestLinkDirPath(catSubDirectoryName)
	linkdirectories.mediaInfo.CreateLinkDirSubDir(catSubDirectoryName)
	linkdirectories.mediaInfo.CreateContentLink()
	linkdirectories.mediaInfo.CreateNewEmptyManifestFile()
	linkdirectories.mediaInfo.CreateManifestLink()

}
