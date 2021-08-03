package mediainformation

import (
	"log"
)

// Creates a directory tree with links based on the parts of the conten source
// file.
func (mediaInfo *MediaInformation) GenerateOldNamePartLinkTree() {

	listOfParts := mediaInfo.ExtractFileNameParts()
	for i := 0; i < len(listOfParts); i++ {

		log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		log.Println("++++++++++++++++++++++++ create old name category ++++++++++++++++++++")
		log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

		mediaInfo.SetAbsoluteContentLinkDirPath("00-olad-name-part/" + listOfParts[i])
		mediaInfo.SetAbsoluteManifestLinkDirPath("00-olad-name-part/" + listOfParts[i])
		mediaInfo.CreateLinkDirSubDir("00-olad-name-part/" + listOfParts[i])
		mediaInfo.CreateContentLink()
		mediaInfo.CreateNewEmptyManifestFile()
		mediaInfo.CreateManifestLink()
	}

}
