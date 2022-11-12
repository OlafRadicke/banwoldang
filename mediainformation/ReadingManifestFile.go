package mediainformation

import (

	cl "github.com/OlafRadicke/banwoldang/customlogger"
	gt "github.com/OlafRadicke/go-gthumb"
)

// ReadingManifestFile Read and pars the manifest file
func (mediaInfo *MediaInformation) ReadingManifestFile() {
	var err error

	mediaInfo.Comments, err = gt.NewCommentsFile(mediaInfo.Comments.FilePath)
	if err != nil {
		cl.ErrorLogger.Println("error to init object: %w", err)
	}


}
