package common

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

//  Reconstruct the path of the conten file from the path of a manifest file.
func reconstructContenFile(mediaInfo *MediaInformation) {
	pathParts := strings.Split(mediaInfo.ManifestFilePath, "/")
	onsUpLevel := len(pathParts) - 2
	mediaInfo.OnsUpPath = strings.Join(pathParts[0:onsUpLevel], "/")

	manifestCoreName := pathParts[len(pathParts)-1]
	manifestExtension := filepath.Ext(manifestCoreName)
	mediaInfo.ContentFileName = manifestCoreName[0 : len(manifestCoreName)-len(manifestExtension)]
	mediaInfo.Extension = filepath.Ext(mediaInfo.ContentFileName)

	mediaInfo.ContentFilePath = mediaInfo.OnsUpPath + "/" + mediaInfo.ContentFileName
	if _, err := os.Stat(mediaInfo.ContentFilePath); err == nil {
		// path/to/whatever exists
		// log.Println("Manifet file and conten file a exit!")
	} else {
		log.Println("File not exist: ", mediaInfo.ContentFilePath)
	}

}
