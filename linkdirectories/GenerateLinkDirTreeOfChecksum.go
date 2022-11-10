package linkdirectories

import (
	cl "github.com/OlafRadicke/banwoldang/customlogger"
	"github.com/OlafRadicke/banwoldang/mediainformation"
)

// GenerateLinkDirTreeOfChecksum Generate directory tree with symlinks to file
// depend of his checksum.
func GenerateLinkDirTreeOfChecksum(mediaInfo *mediainformation.MediaInformation) {

	var (
		firstChar string = ""
	)
	cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	cl.InfoLogger.Println("++++++++++++++++++++++ Create checksum link ++++++++++++++++++++++")
	cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	checksum := mediaInfo.HashValue
	if len(checksum) < 1 {
		cl.ErrorLogger.Println("Is this checksum relay correct?: ", checksum)
		firstChar = "unknow"
	} else {
		firstChar = string(checksum[0])
	}
	catSubDirectoryName := "00-checksum/" + firstChar
	cl.InfoLogger.Println("First char: ", firstChar, " of ", checksum)
	cl.InfoLogger.Println("Use: ", catSubDirectoryName)
	mediaInfo.SetAbsoluteContentLinkDirPath(catSubDirectoryName)
	mediaInfo.SetAbsoluteManifestLinkDirPath(catSubDirectoryName)
	mediaInfo.CreateLinkDirSubDir(catSubDirectoryName)
	err := mediaInfo.CreateContentLink()
	if err != nil {
		cl.ErrorLogger.Println("Switch from hash to file name: ", mediaInfo.ContentFileName)
		cl.DuplicateLogger.Println(err)
		catSubDirectoryName := "00-checksum/duplicates"
		mediaInfo.SetHashValue(mediaInfo.ContentFileName)
		mediaInfo.SetAbsoluteContentLinkDirPath(catSubDirectoryName)
		mediaInfo.SetAbsoluteManifestLinkDirPath(catSubDirectoryName)
		mediaInfo.CreateLinkDirSubDir(catSubDirectoryName)
	}
	mediaInfo.CreateNewEmptyManifestFile()
	mediaInfo.CreateManifestLink()

}
