package mediainformation

import cl "github.com/OlafRadicke/banwoldang/customlogger"

// Creates a directory tree with links based on the parts of the conten source
// file.
func (mediaInfo *MediaInformation) GenerateOldNamePartLinkTree() {

	listOfParts := mediaInfo.ExtractFileNameParts()
	for i := 0; i < len(listOfParts); i++ {

		cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		cl.InfoLogger.Println("++++++++++++++++++++++++ create old name category ++++++++++++++++++++")
		cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

		mediaInfo.SetAbsoluteContentLinkDirPath("00-olad-name-part/" + listOfParts[i])
		mediaInfo.SetAbsoluteManifestLinkDirPath("00-olad-name-part/" + listOfParts[i])
		mediaInfo.CreateLinkDirSubDir("00-olad-name-part/" + listOfParts[i])
		mediaInfo.CreateContentLink()
		mediaInfo.CreateNewEmptyManifestFile()
		mediaInfo.CreateManifestLink()
	}

}
