package common

import (
	"log"
	"os"
	"path/filepath"
)

type filetree interface {
	GoThroughCollection()
	PrintAll()
	fileHandler(path string, info os.FileInfo, err error) error
}

type FileTree struct {
	// The start path for searching media files
	StartPath string
	// Location for the generic directory
	GenericDir string
	// The number of founded manifest Files
	Findings int
}

type MediaInformation struct {
	// Relative path of manifest xml file.
	ManifestPath string
	// The relative directory path on up with the media files.
	OnsUpPath string
	// The name of the media file
	ContentFileName string
	// The media file with the relative path
	ContentFilePath string
	// The hash value of the media file
	HashValue string
	// The extension of the media file
	Extension string
	// The relative path for the generic file tree
	GenFilePath string
	// The list with the categories of a media file
	Categories []string
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
	// log.Println(searchPath, info.Name(), info.Size())
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
			createMediaFileHash(&mediaInfo)
			genericFileTree(&mediaInfo, fileTree)
		}

	}
	return nil
}
