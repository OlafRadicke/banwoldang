package mediainformation

import (

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// SaveManifestFile Saved the manifest file
func (mediaInfo *MediaInformation) SaveManifestFile() {
	var err error
	cl.InfoLogger.Println("Save Manifest: ", mediaInfo.Comments.FilePath)
	err = mediaInfo.Comments.Save()
	if err != nil {
		cl.ErrorLogger.Println(err)
	}
}
