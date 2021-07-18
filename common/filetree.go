package common

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type filetree interface {
	GoThroughCollection()
	PrintAll()
	fileHandler(path string, info os.FileInfo, err error) error
}

type FileTree struct {
	StartPath string
	Findings  int
}

type MediaInformation struct {
	ManifestPath    string
	OnsUpPath       string
	ContentFileName string
	ContentFilePath string
}

func (fileTree *FileTree) GoThroughCollection() {
	err := filepath.Walk(fileTree.StartPath, fileTree.fileHandler)
	if err != nil {
		log.Println(err)
	}
}

func (fileTree *FileTree) fileHandler(searchPath string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	// log.Println(path, info.Name(), info.Size())
	if info.IsDir() {
		// log.Println("Just a directory")
	} else {
		fileTree.Findings++
		mediaInfo := MediaInformation{}
		if filepath.Ext(info.Name()) == ".xml" {
			// log.Println("===========================================")
			// log.Println("Manifest file: ", searchPath)
			// log.Println("Content file: ", reconstructContenFile(searchPath))
			mediaInfo.ManifestPath = searchPath
			reconstructContenFile(&mediaInfo)
			readingManifestFile(&mediaInfo)
		}

	}
	return nil
}

//  Reconstruct the path of the conten file from the path of a manifest file.
func reconstructContenFile(mediaInfo *MediaInformation) {
	pathParts := strings.Split(mediaInfo.ManifestPath, "/")
	onsUpLevel := len(pathParts) - 2
	mediaInfo.OnsUpPath = strings.Join(pathParts[0:onsUpLevel], "/")

	manifestCoreName := pathParts[len(pathParts)-1]
	extension := filepath.Ext(manifestCoreName)
	mediaInfo.ContentFileName = manifestCoreName[0 : len(manifestCoreName)-len(extension)]

	mediaInfo.ContentFilePath = mediaInfo.OnsUpPath + "/" + mediaInfo.ContentFileName
	if _, err := os.Stat(mediaInfo.ContentFilePath); err == nil {
		// path/to/whatever exists
		// log.Println("Manifet file and conten file a exit!")
	} else {
		log.Println("File not exist: ", mediaInfo.ContentFilePath)
	}

}
