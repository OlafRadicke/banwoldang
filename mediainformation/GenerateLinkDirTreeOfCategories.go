package mediainformation

import (
	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// GenerateLinkDirTreeOfCategories Creates a directory tree with links based on the
// categories in the xml manifest files.
func (mediaInfo *MediaInformation) GenerateLinkDirTreeOfCategories() {

	if len(mediaInfo.Categories) > 0 {
		cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		cl.InfoLogger.Println("+++++++++++++++++++ create links for ", len(mediaInfo.Categories), " categories ++++++++++++++++++++")
		cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	}

	for i := 0; i < len(mediaInfo.Categories); i++ {
		if len(mediaInfo.Categories) < 6 {
			mediaInfo.GenerateLinkDirTreeOfCategoryCount()
		}

		mediaInfo.SetAbsoluteContentLinkDirPath("categories/" + mediaInfo.Categories[i])
		mediaInfo.SetAbsoluteManifestLinkDirPath("categories/" + mediaInfo.Categories[i])
		mediaInfo.CreateLinkDirSubDir("categories/" + mediaInfo.Categories[i])
		mediaInfo.CreateContentLink()
		mediaInfo.CreateNewEmptyManifestFile()
		mediaInfo.CreateManifestLink()

	}
}
