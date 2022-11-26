package linkdirectories

import (
	cl "github.com/OlafRadicke/banwoldang/customlogger"
	"github.com/OlafRadicke/banwoldang/mediainformation"
)

// GenerateLinkDirTreeOfChecksum Generate directory tree with symlinks to file
// depend of his checksum.
func GenerateLinkDirTreeOfChecksum(mediaInfo *mediainformation.MediaInformation) (error) {
	var (
		firstChar string = ""
		err error
	)
	cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	cl.InfoLogger.Println("++++++++++++++++++++++ Create checksum link ++++++++++++++++++++++")
	cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	checksum, err := mediaInfo.GetHashValue()
	if err != nil {
		cl.ErrorLogger.Println(err)
		cl.DuplicateLogger.Println(mediaInfo.AbsoluteContentSourcePath)
		catSubDirectoryName := "00-checksum/00-duplicates"
		mediaInfo.SetAbsoluteContentLinkDirPath(catSubDirectoryName)
		mediaInfo.SetAbsoluteManifestLinkDirPath(catSubDirectoryName)
		mediaInfo.CreateLinkDirSubDir(catSubDirectoryName)
		mediaInfo.Comments.AddCategory("00-dublette")
		return err
	}
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
	err = mediaInfo.CreateContentLink()
	if err != nil {
		cl.DuplicateLogger.Println(mediaInfo.AbsoluteContentSourcePath)
		catSubDirectoryName := "00-checksum/00-duplicates"
		mediaInfo.SetAbsoluteContentLinkDirPath(catSubDirectoryName)
		mediaInfo.SetAbsoluteManifestLinkDirPath(catSubDirectoryName)
		mediaInfo.CreateLinkDirSubDir(catSubDirectoryName)
		mediaInfo.Comments.AddCategory("00-dublette")
		err = mediaInfo.SaveManifestFile()
		if err != nil {
			cl.ErrorLogger.Println(err)
		}
	}
	mediaInfo.CreateManifestLink()
	return nil
}
