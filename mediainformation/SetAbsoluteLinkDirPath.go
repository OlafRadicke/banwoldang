package mediainformation

import (
	"path/filepath"
    "fmt"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// SetAbsoluteLinkDirPath Create the absolute link to the directory with the
// link tree.
// @linkDir param is the value with an relative path to the link tree directory.
func (mediaInfo *MediaInformation) SetAbsoluteLinkDirPath(linkDir string) {
	prefix := "/link-tree"
	absolutLinkDir, err1 := filepath.Abs(linkDir + prefix)
	if err1 != nil {
		fmt.Println("error! check the log files for more information")
		cl.ErrorLogger.Fatal(err1)
	}
	mediaInfo.AbsoluteLinkDirPath = absolutLinkDir
}
