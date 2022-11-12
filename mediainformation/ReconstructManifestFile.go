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
	cl.InfoLogger.Println("mediaInfo.ContentFileName: ", mediaInfo.ContentFileName)
	cl.InfoLogger.Println("baseDir: ", baseDir)
	reconstructedpath := baseDir + "/.comments/" + mediaInfo.ContentFileName + ".xml"
	cl.InfoLogger.Println("reconstructedpath: ", reconstructedpath)
	_, err = os.Stat(reconstructedpath)
	if err != nil {
		mediaInfo.Comments, err = gt.NewCommentsFile(reconstructedpath)
		if err != nil {
			cl.InfoLogger.Println("error to init object: ", err)
		}

		cl.InfoLogger.Println("Comment file not found. Create new: ", reconstructedpath)
		mediaInfo.Comments.AddCategory("00-script-create-manifest")
		mediaInfo.SaveManifestFile()
	}

}
