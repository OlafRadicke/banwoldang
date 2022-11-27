package linkdirectories

import (
	cl "github.com/OlafRadicke/banwoldang/customlogger"
	"github.com/OlafRadicke/banwoldang/mediainformation"
	"github.com/OlafRadicke/banwoldang/statistics"
)

// GenerateLinkDirTreeOfChecksum Generate directory tree with symlinks to file
// depend of his checksum.
func GenerateLinkDirTreeOfChecksum(mediaInfo *mediainformation.MediaInformation, statistics *statistics.Statistics) (error) {
	var (
		firstChar string = ""
		err error
		checksum string
	)
	cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	cl.InfoLogger.Println("++++++++++++++++++++++ Create checksum link ++++++++++++++++++++++")
	cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	checksum, err = mediaInfo.GetHashValue()
	if err != nil {
		cl.ErrorLogger.Println(err)
		return err
	}
	locationsCount, err := statistics.GetCheckSumLocationsCount(checksum)
	if err != nil {
		cl.ErrorLogger.Println(err)
		return err
	} else{
		if locationsCount > 1 {
			mediaInfo.Comments.AddCategory("00-dublette")
		}
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

	mediaInfo.CreateManifestLink()
	return nil
}
