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
	StartPath string
	Findings  int
}

type MediaInformation struct {
	ManifestPath    string
	OnsUpPath       string
	ContentFileName string
	ContentFilePath string
	HashValue       string
	Extension       string
	GenFilePath     string
	Categories      []string
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
			createMediaFileHash(&mediaInfo)
			genericFileTree(&mediaInfo)
		}

	}
	return nil
}
