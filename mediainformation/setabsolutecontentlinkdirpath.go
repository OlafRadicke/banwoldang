package mediainformation

import (
	"path/filepath"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// Set or create the absolute link path of the contnet file for the link directory tree.
// @subDir param is the name of a sub directory for link directory tree.
func (mediaInfo *MediaInformation) SetAbsoluteContentLinkDirPath(subDir string) {
	genFilePath := mediaInfo.AbsoluteLinkDirPath + "/" + subDir + "/" + mediaInfo.HashValue + mediaInfo.Extension
	absolutLinkTarget, err1 := filepath.Abs(genFilePath)
	if err1 != nil {
		cl.ErrorLogger.Fatal(err1)
	}
	mediaInfo.AbsoluteContentLinkDirPath = absolutLinkTarget
}
