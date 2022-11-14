package mediainformation

import (
	"errors"
	"os"
	"path/filepath"
	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// SaveManifestFile Saved the manifest file
func (mediaInfo *MediaInformation) SaveManifestFile() (error){
	var err error
	if filepath.Ext(mediaInfo.Comments.FilePath) != ".xml" {
		// cl.ErrorLogger.Fatal("Thats a not a right name for an command file: ", mediaInfo.Comments.FilePath)
		return errors.New("Thats a not a right name for an command file: " + mediaInfo.Comments.FilePath)
	}
	commentDir := filepath.Dir(mediaInfo.Comments.FilePath)
	if _, err = os.Stat(commentDir); os.IsNotExist(err) {
		if err = os.Mkdir(commentDir, os.ModePerm); err != nil {
			cl.ErrorLogger.Println(err)
		}
	}
	cl.InfoLogger.Println("Save Manifest: ", mediaInfo.Comments.FilePath)
	err = mediaInfo.Comments.Save()
	if err != nil {
		cl.ErrorLogger.Println(err)
		return err
	}
	return nil
}
