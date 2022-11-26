package mediainformation

import (
	"path/filepath"
    "fmt"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// SetAbsoluteContentLinkDirPath Set or create the absolute link path of the
// contnet file for the link directory tree.
// @subDir param is the name of a sub directory for link directory tree.
func (mediaInfo *MediaInformation) SetAbsoluteContentLinkDirPath(subDir string) (error) {
	var err error
	checksum, err := mediaInfo.GetHashValue()
	if err != nil {
		cl.ErrorLogger.Println(err)
		return err
	}
	mediaInfo.linkSubDir = subDir
	genFilePath := mediaInfo.AbsoluteLinkDirPath + "/" + mediaInfo.linkSubDir + "/" + checksum + mediaInfo.Extension
	absolutLinkTarget, err := filepath.Abs(genFilePath)
	if err != nil {
		fmt.Println("error! check the log files for more information")
		cl.ErrorLogger.Fatal(err)
	}
	mediaInfo.AbsoluteContentLinkDirPath = absolutLinkTarget
	return nil
}
