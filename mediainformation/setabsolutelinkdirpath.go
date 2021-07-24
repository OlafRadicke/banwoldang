package mediainformation

import (
	"log"
	"path/filepath"
)

// Create the absolute link to the directory with the link tree.
// @linkDir param is the value with an relative path to the link tree directory.
func (mediaInfo *MediaInformation) SetAbsoluteLinkDirPath(linkDir string) {
	prefix := "/link-tree"
	absolutLinkDir, err1 := filepath.Abs(linkDir + prefix)
	if err1 != nil {
		log.Fatal(err1)
	}
	mediaInfo.AbsoluteLinkDirPath = absolutLinkDir
}
