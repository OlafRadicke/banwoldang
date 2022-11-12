package linkdirectories

import (
	cl "github.com/OlafRadicke/banwoldang/customlogger"
	"github.com/OlafRadicke/banwoldang/mediainformation"
)

// GenerateLinkDirTreeOfCategories Creates a directory tree with links based on the
// categories in the xml manifest files.
func GenerateLinkDirTreeOfCategories(mediaInfo *mediainformation.MediaInformation) {
	if len(mediaInfo.Comments.GetCategories()) > 0 {
		cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		cl.InfoLogger.Println("+++++++++++++++++++ create links for ", len(mediaInfo.Comments.GetCategories()), " categories ++++++++++++++++++++")
		cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	}

	for _, category := range mediaInfo.Comments.GetCategories(){

		mediaInfo.SetAbsoluteContentLinkDirPath("categories/" + category.Value)
		mediaInfo.SetAbsoluteManifestLinkDirPath("categories/" + category.Value)
		mediaInfo.CreateLinkDirSubDir("categories/" + category.Value)
		mediaInfo.CreateContentLink()
		// mediaInfo.CreateNewEmptyManifestFile()
		mediaInfo.CreateManifestLink()

	}

	if len(mediaInfo.Comments.GetCategories()) < 6 {
		GenerateLinkDirTreeOfCategoryCount(mediaInfo)
	}
}
