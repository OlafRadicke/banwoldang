package mediainformation

import (
	"os"
	"path/filepath"
	"strings"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// ReconstructContenSourceFile Reconstruct the source path of the conten file
// based on the path of a manifest file.
func (mediaInfo *MediaInformation) ReconstructContenSourceFile() error {
	pathParts := strings.Split(mediaInfo.Comments.FilePath, "/")
	onsUpLevel := len(pathParts) - 2
	mediaInfo.OnsUpPath = strings.Join(pathParts[0:onsUpLevel], "/")

	manifestCoreName := pathParts[len(pathParts)-1]
	manifestExtension := filepath.Ext(manifestCoreName)
	mediaInfo.ContentFileName = manifestCoreName[0 : len(manifestCoreName)-len(manifestExtension)]
	mediaInfo.Extension = filepath.Ext(mediaInfo.ContentFileName)

	mediaInfo.AbsoluteContentSourcePath = mediaInfo.OnsUpPath + "/" + mediaInfo.ContentFileName
	if _, err := os.Stat(mediaInfo.AbsoluteContentSourcePath); err == nil {
		// path/to/whatever exists
		// cl.InfoLogger.Println("Manifet file and conten file a exit!")
		return nil
	} else {
		cl.ErrorLogger.Println("Conten file of manifest file not exist: ", mediaInfo.AbsoluteContentSourcePath)
		cl.OrphanLogger.Println(mediaInfo.Comments.FilePath)
		return err
	}

}
