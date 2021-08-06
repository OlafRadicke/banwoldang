package mediainformation

import (
	"os"
	"path/filepath"
	"strings"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

//  Reconstruct the path of the conten file from the path of a manifest file.
func (mediaInfo *MediaInformation) ReconstructManifestFile() {
	pathParts := strings.Split(mediaInfo.AbsoluteContentSourcePath, "/")
	mediaInfo.ContentFileName = pathParts[len(pathParts)-1]
	baseDir := filepath.Dir(mediaInfo.AbsoluteContentSourcePath)
	mediaInfo.Extension = filepath.Ext(mediaInfo.ContentFileName)
	mediaInfo.AbsoluteManifestSourcePath = baseDir + "/.comments/" + mediaInfo.ContentFileName + ".xml"

	if _, err := os.Stat(mediaInfo.AbsoluteContentSourcePath); err == nil {
		// path/to/whatever exists
		// cl.InfoLogger.Println("Manifet file and conten file a exit!")
	} else {
		cl.InfoLogger.Println("Manifest file not exist: ", mediaInfo.AbsoluteManifestSourcePath)
	}

}
