package mediainformation

import (

	cl "github.com/OlafRadicke/banwoldang/customlogger"
	// gt "github.com/OlafRadicke/go-gthumb"
	"fmt"
)

// SaveManifestFile Saved the manifest file
func (mediaInfo *MediaInformation) SaveManifestFile() {
	var err error

	fmt.Println("Save MAnifest: ", mediaInfo.Comments.FilePath)

	err = mediaInfo.Comments.Save()
	if err != nil {
		cl.ErrorLogger.Println(err)
	}
}
