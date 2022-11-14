package mediainformation

import (
	"path/filepath"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// SetAbsoluteManifestLinkDirPath Create the absolute link path of the manifest
// file for the link directory tree.
// @subDir param is the name of a sub directory for link directory tree.
func (mediaInfo *MediaInformation) SetAbsoluteManifestLinkDirPath(subDir string) {
	checksum, err := mediaInfo.GetHashValue()
	if err != nil {
		cl.ErrorLogger.Println(err)
		return
	}

	genFilePath := mediaInfo.AbsoluteLinkDirPath + "/" + subDir + "/.comments/" + checksum + mediaInfo.Extension + ".xml"
	absolutLinkTarget, err1 := filepath.Abs(genFilePath)
	if err1 != nil {
		cl.ErrorLogger.Fatal(err1)
	}
	mediaInfo.AbsoluteManifestLinkDirPath = absolutLinkTarget
}
