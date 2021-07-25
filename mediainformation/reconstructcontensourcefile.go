package mediainformation

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Reconstruct the source path of the conten file based on the path of a
// manifest file.
func (mediaInfo *MediaInformation) ReconstructContenSourceFile() {
	log.Println("mediaInfo.AbsoluteManifestSourcePath: ", mediaInfo.AbsoluteManifestSourcePath)
	pathParts := strings.Split(mediaInfo.AbsoluteManifestSourcePath, "/")
	log.Println("pathParts: ", pathParts)
	onsUpLevel := len(pathParts) - 2
	mediaInfo.OnsUpPath = strings.Join(pathParts[0:onsUpLevel], "/")

	manifestCoreName := pathParts[len(pathParts)-1]
	manifestExtension := filepath.Ext(manifestCoreName)
	mediaInfo.ContentFileName = manifestCoreName[0 : len(manifestCoreName)-len(manifestExtension)]
	mediaInfo.Extension = filepath.Ext(mediaInfo.ContentFileName)

	mediaInfo.AbsoluteContentSourcePath = mediaInfo.OnsUpPath + "/" + mediaInfo.ContentFileName
	if _, err := os.Stat(mediaInfo.AbsoluteContentSourcePath); err == nil {
		// path/to/whatever exists
		// log.Println("Manifet file and conten file a exit!")
	} else {
		log.Println("File not exist: ", mediaInfo.AbsoluteContentSourcePath)
	}

}
