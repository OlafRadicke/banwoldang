package mediainformation

import (
	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// GenerateLinkDirTreeOfChecksum Generate directory tree with symlinks to file
// depend of his checksum.
func (mediaInfo *MediaInformation) GenerateLinkDirTreeOfChecksum() {

	cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	cl.InfoLogger.Println("++++++++++++++++++++++ Create checksum link ++++++++++++++++++++++")
	cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	checksum := mediaInfo.HashValue
	firstChar := string(checksum[0])
	catSubDirectoryName := "00-checksum/" + firstChar
	cl.InfoLogger.Println("First char: ", firstChar, " of ", checksum)
	cl.InfoLogger.Println("Use: ", catSubDirectoryName)
	mediaInfo.SetAbsoluteContentLinkDirPath(catSubDirectoryName)
	mediaInfo.SetAbsoluteManifestLinkDirPath(catSubDirectoryName)
	mediaInfo.CreateLinkDirSubDir(catSubDirectoryName)
	mediaInfo.CreateContentLink()
	mediaInfo.CreateNewEmptyManifestFile()
	mediaInfo.CreateManifestLink()

}
