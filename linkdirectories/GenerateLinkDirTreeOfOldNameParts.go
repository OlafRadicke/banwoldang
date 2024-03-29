package linkdirectories

import (
	cl "github.com/OlafRadicke/banwoldang/customlogger"
	"github.com/OlafRadicke/banwoldang/mediainformation"
)

// GenerateLinkDirTreeOfOldNameParts Creates a directory tree with links based on the
// parts of the conten source file.
func GenerateLinkDirTreeOfOldNameParts(mediaInfo *mediainformation.MediaInformation) {

	listOfParts := mediaInfo.ExtractFileNameParts()

	if len(listOfParts) > 0 {
		cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		cl.InfoLogger.Println("++++++++++++++++++++++++ create old name category ++++++++++++++++++++")
		cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	}

	for i := 0; i < len(listOfParts); i++ {
		var namePart string
		mediaInfo.SetAbsoluteContentLinkDirPath("00-old-name-parts/" + listOfParts[i])
		mediaInfo.SetAbsoluteManifestLinkDirPath("00-old-name-parts/" + listOfParts[i])
		mediaInfo.CreateLinkDirSubDir("00-old-name-parts/" + listOfParts[i])
		mediaInfo.CreateContentLink()
		// mediaInfo.CreateNewEmptyManifestFile()
		mediaInfo.CreateManifestLink()
		namePart = listOfParts[i]
		mediaInfo.Statistics.ContPartsOfNames(namePart)
	}

}
