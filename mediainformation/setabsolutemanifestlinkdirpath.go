package mediainformation

import (
	"log"
	"path/filepath"
)

// Create the absolute link path of the manifest file for the link directory tree.
// @subDir param is the name of a sub directory for link directory tree.
func (mediaInfo *MediaInformation) SetAbsoluteManifestLinkDirPath(subDir string) {
	genFilePath := mediaInfo.AbsoluteLinkDirPath + "/" + subDir + "/.comments/" + mediaInfo.HashValue + mediaInfo.Extension + ".xml"
	absolutLinkTarget, err1 := filepath.Abs(genFilePath)
	if err1 != nil {
		log.Fatal(err1)
	}
	log.Println("SetAbsoluteManifestLinkDirPath: ", absolutLinkTarget)
	mediaInfo.AbsoluteManifestLinkDirPath = absolutLinkTarget
}
