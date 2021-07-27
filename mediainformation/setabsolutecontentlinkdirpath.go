package mediainformation

import (
	"log"
	"path/filepath"
)

// Set or create the absolute link path of the contnet file for the link directory tree.
// @subDir param is the name of a sub directory for link directory tree.
func (mediaInfo *MediaInformation) SetAbsoluteContentLinkDirPath(subDir string) {
	genFilePath := mediaInfo.AbsoluteLinkDirPath + "/" + subDir + "/" + mediaInfo.HashValue + mediaInfo.Extension
	absolutLinkTarget, err1 := filepath.Abs(genFilePath)
	if err1 != nil {
		log.Fatal(err1)
	}
	mediaInfo.AbsoluteContentLinkDirPath = absolutLinkTarget
}
