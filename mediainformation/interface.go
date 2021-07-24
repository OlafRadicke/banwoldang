package mediainformation

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type mediainformation interface {
	CreateMediaFileHash()
	GenericSingleCatFileTree(string) // TODO
	GenericFileTree(string)          // TODO
	GenericNonCatFileTree(string)    // TODO
	ReadingManifestFile()
	ReconstructContenFile()
	ReconstructManifestFile()
}

type MediaInformation struct {
	// Relative path of manifest xml file.
	ManifestFilePath string
	// The relative directory path on up with the media files.
	OnsUpPath string
	// The name of the media file
	ContentFileName string
	// The source path of the media as an absolute path
	AbsoluteContentSourcePath string
	// Absolute path for the target link in the link directory rom the conten file.
	AbsoluteContentLinkDirPath string
	// The hash value of the media file
	HashValue string
	// The extension of the media file
	Extension string
	// The relative path for the generic file tree
	GenFilePath string
	// The list with the categories of a media file
	Categories []string
}

//  Reconstruct the path of the conten file from the path of a manifest file.
func (mediaInfo *MediaInformation) ReconstructContenFile() {
	log.Println("mediaInfo.ManifestFilePath: ", mediaInfo.ManifestFilePath)
	pathParts := strings.Split(mediaInfo.ManifestFilePath, "/")
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
