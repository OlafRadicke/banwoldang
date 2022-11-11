package linkdirectories

import (
	cl "github.com/OlafRadicke/banwoldang/customlogger"
	// "github.com/OlafRadicke/banwoldang/mediainformation"
)

// GenerateLinkDirTreeOfChecksum Generate directory tree with symlinks to file
// depend of his checksum.
func (linkdirectories  *Linkdirectories) GenerateLinkDirTreeOfChecksum() {

	var (
		firstChar string = ""
	)
	cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	cl.InfoLogger.Println("++++++++++++++++++++++ Create checksum link ++++++++++++++++++++++")
	cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	checksum := linkdirectories.mediaInfo.HashValue
	if len(checksum) < 1 {
		cl.ErrorLogger.Println("Is this checksum relay correct?: ", checksum)
		firstChar = "unknow"
	} else {
		firstChar = string(checksum[0])
	}
	catSubDirectoryName := "00-checksum/" + firstChar
	cl.InfoLogger.Println("First char: ", firstChar, " of ", checksum)
	cl.InfoLogger.Println("Use: ", catSubDirectoryName)
	linkdirectories.mediaInfo.SetAbsoluteContentLinkDirPath(catSubDirectoryName)
	linkdirectories.mediaInfo.SetAbsoluteManifestLinkDirPath(catSubDirectoryName)
	linkdirectories.mediaInfo.CreateLinkDirSubDir(catSubDirectoryName)
	err := linkdirectories.mediaInfo.CreateContentLink()
	if err != nil {
		cl.ErrorLogger.Println("Switch from hash to file name: ", linkdirectories.mediaInfo.ContentFileName)
		cl.DuplicateLogger.Println(err)
		catSubDirectoryName := "00-checksum/duplicates"
		linkdirectories.mediaInfo.SetHashValue(linkdirectories.mediaInfo.ContentFileName)
		linkdirectories.mediaInfo.SetAbsoluteContentLinkDirPath(catSubDirectoryName)
		linkdirectories.mediaInfo.SetAbsoluteManifestLinkDirPath(catSubDirectoryName)
		linkdirectories.mediaInfo.CreateLinkDirSubDir(catSubDirectoryName)
	}
	// linkdirectories.mediaInfo.CreateNewEmptyManifestFile()
	linkdirectories.mediaInfo.CreateManifestLink()

}
