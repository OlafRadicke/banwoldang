package linkdirectories

import (
	cl "github.com/OlafRadicke/banwoldang/customlogger"
	// "github.com/OlafRadicke/banwoldang/mediainformation"
	"fmt"
)

// GenerateLinkDirTreeOfCategories Creates a directory tree with links based on the
// categories in the xml manifest files.
func (linkdirectories  *Linkdirectories)GenerateLinkDirTreeOfCategories() {
	fmt.Println("GenerateLinkDirTreeOfCategories \n")
	if len(linkdirectories.getCategories()) > 0 {
		cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		cl.InfoLogger.Println("+++++++++++++++++++ create links for ", len(linkdirectories.getCategories()), " categories ++++++++++++++++++++")
		cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	}

	fmt.Println("len(linkdirectories.getCategories())", len(linkdirectories.getCategories()), " \n")

	for _, category := range linkdirectories.getCategories(){

		linkdirectories.mediaInfo.SetAbsoluteContentLinkDirPath("categories/" + category.Value)
		linkdirectories.mediaInfo.SetAbsoluteManifestLinkDirPath("categories/" + category.Value)
		linkdirectories.mediaInfo.CreateLinkDirSubDir("categories/" + category.Value)
		linkdirectories.mediaInfo.CreateContentLink()
		linkdirectories.mediaInfo.CreateNewEmptyManifestFile()
		linkdirectories.mediaInfo.CreateManifestLink()

	}

	if len(linkdirectories.getCategories()) < 6 {
		linkdirectories.GenerateLinkDirTreeOfCategoryCount()
	}
}
