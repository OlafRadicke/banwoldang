package mediainformation

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

//  Reconstruct the path of the conten file from the path of a manifest file.
func (mediaInfo *MediaInformation) ReconstructManifestFile() {
	pathParts := strings.Split(mediaInfo.AbsoluteContentSourcePath, "/")
	mediaInfo.ContentFileName = pathParts[len(pathParts)-1]
	baseDir := filepath.Dir(mediaInfo.AbsoluteContentSourcePath)
	mediaInfo.Extension = filepath.Ext(mediaInfo.ContentFileName)
	mediaInfo.ManifestFilePath = baseDir + ".comments/" + mediaInfo.ContentFileName + ".xml"

	if _, err := os.Stat(mediaInfo.AbsoluteContentSourcePath); err == nil {
		// path/to/whatever exists
		// log.Println("Manifet file and conten file a exit!")
	} else {
		log.Println("Manifest file not exist: ", mediaInfo.ManifestFilePath)
	}

}
