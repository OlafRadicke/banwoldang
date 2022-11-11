package mediainformation

import (
	"os"
	"path/filepath"
	"strings"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
	gt "github.com/OlafRadicke/go-gthumb"
)

//  ReconstructManifestFile Reconstruct the path of the conten file from the path of a manifest file.
func (mediaInfo *MediaInformation) ReconstructManifestFile() {
	var err error
	pathParts := strings.Split(mediaInfo.AbsoluteContentSourcePath, "/")
	mediaInfo.ContentFileName = pathParts[len(pathParts)-1]
	baseDir := filepath.Dir(mediaInfo.AbsoluteContentSourcePath)
	mediaInfo.Extension = filepath.Ext(mediaInfo.ContentFileName)
	mediaInfo.AbsoluteManifestSourcePath = baseDir + "/.comments/" + mediaInfo.ContentFileName + ".xml"

	if _, err = os.Stat(mediaInfo.AbsoluteContentSourcePath); err == nil {
		cl.InfoLogger.Println("Manifet file and conten file a exit!")
		mediaInfo.Comments, err = gt.NewCommentsFile(mediaInfo.AbsoluteContentSourcePath)
		if err != nil {
			cl.InfoLogger.Println("error to init object: ", err)
		}
	} else {
		cl.InfoLogger.Println("Manifest file not exist: ", mediaInfo.AbsoluteManifestSourcePath)
		mediaInfo.Comments, err = gt.NewCommentsFile(mediaInfo.AbsoluteContentSourcePath)
		if err != nil {
			cl.InfoLogger.Println("error to init object: ", err)
		}
		mediaInfo.Comments.AddCategory("00-script-create-manifest")
		mediaInfo.Comments.Save()
	}


}
