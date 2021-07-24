package mediainformation

import (
	"log"
	"path/filepath"
)

// Create the absolute link path of the contnet file for the link directory tree.
// @genericDir param is the base directory of the link directory tree.
// @subDir param is the name of a sub directory for link directory tree.
func (mediaInfo *MediaInformation) SetAbsoluteContentLinkDirPath(genericDir string, subDir string) {
	genFilePath := genericDir + "/gereric-tree/" + subDir + "/" + mediaInfo.HashValue + mediaInfo.Extension
	absolutLinkTarget, err1 := filepath.Abs(genFilePath)
	if err1 != nil {
		log.Fatal(err1)
	}
	log.Println("absolutLinkTarget: ", absolutLinkTarget)
	mediaInfo.AbsoluteContentLinkDirPath = absolutLinkTarget
}
