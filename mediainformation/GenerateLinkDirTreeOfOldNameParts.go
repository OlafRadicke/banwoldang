package mediainformation

import cl "github.com/OlafRadicke/banwoldang/customlogger"

// GenerateLinkDirTreeOfOldNameParts Creates a directory tree with links based on the
// parts of the conten source file.
func (mediaInfo *MediaInformation) GenerateLinkDirTreeOfOldNameParts() {

	listOfParts := mediaInfo.ExtractFileNameParts()

	if len(listOfParts) > 0 {
		cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		cl.InfoLogger.Println("++++++++++++++++++++++++ create old name category ++++++++++++++++++++")
		cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	}

	for i := 0; i < len(listOfParts); i++ {
		mediaInfo.SetAbsoluteContentLinkDirPath("00-old-name-parts/" + listOfParts[i])
		mediaInfo.SetAbsoluteManifestLinkDirPath("00-old-name-parts/" + listOfParts[i])
		mediaInfo.CreateLinkDirSubDir("00-old-name-parts/" + listOfParts[i])
		mediaInfo.CreateContentLink()
		mediaInfo.CreateNewEmptyManifestFile()
		mediaInfo.CreateManifestLink()
		mediaInfo.Statistics.ContPartsOfNames(listOfParts[i])
	}

}
