package mediainformation

import (
	"path/filepath"
    "fmt"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// SetAbsoluteManifestLinkDirPath Create the absolute link path of the manifest
// file for the link directory tree.
// @subDir param is the name of a sub directory for link directory tree.
func (mediaInfo *MediaInformation) SetAbsoluteManifestLinkDirPath(subDir string) {
	var err error
	checksum, err := mediaInfo.GetHashValue()
	if err != nil {
		cl.ErrorLogger.Println(err)
		return
	}

	genFilePath := mediaInfo.AbsoluteLinkDirPath + "/" + subDir + "/.comments/" + checksum + mediaInfo.Extension + ".xml"
	absolutLinkTarget, err := filepath.Abs(genFilePath)
	if err != nil {
		fmt.Println("error! check the log files for more information")
		cl.ErrorLogger.Fatal(err)
	}
	mediaInfo.AbsoluteManifestLinkDirPath = absolutLinkTarget
}
