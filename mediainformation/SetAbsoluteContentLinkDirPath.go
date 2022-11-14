package mediainformation

import (
	"path/filepath"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// SetAbsoluteContentLinkDirPath Set or create the absolute link path of the
// contnet file for the link directory tree.
// @subDir param is the name of a sub directory for link directory tree.
func (mediaInfo *MediaInformation) SetAbsoluteContentLinkDirPath(subDir string) {
	checksum, err := mediaInfo.GetHashValue()
	if err != nil {
		cl.ErrorLogger.Println(err)
		return
	}
	genFilePath := mediaInfo.AbsoluteLinkDirPath + "/" + subDir + "/" + checksum + mediaInfo.Extension
	absolutLinkTarget, err1 := filepath.Abs(genFilePath)
	if err1 != nil {
		cl.ErrorLogger.Fatal(err1)
	}
	mediaInfo.AbsoluteContentLinkDirPath = absolutLinkTarget
}
