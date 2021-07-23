package common

import (
	"log"
	"os"
	"path/filepath"

	"github.com/OlafRadicke/banwoldang/mediainformation"
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
	log.Println("func params:", searchPath, info.Name(), info.Size())
	if info.IsDir() {
		// log.Println("Just a directory")
	} else {
		fileTree.Findings++
		mediaInfo := mediainformation.MediaInformation{}
		if filepath.Ext(info.Name()) == ".xml" {

			log.Println("===========================================")
			log.Println("searchPath: ", searchPath)
			mediaInfo.ManifestFilePath = searchPath
			log.Println("Manifest file: ", searchPath)
			log.Println("mediaInfo.ManifestFilePath: ", mediaInfo.ManifestFilePath)
			log.Println("===========================================")

			mediaInfo.ReconstructContenFile()
			mediaInfo.ReadingManifestFile()
			mediaInfo.CreateMediaFileHash()
			genericFileTree(&mediaInfo, fileTree)
		} else {
			log.Println("-------------------------------------------")
			log.Println("searchPath: ", searchPath)
			mediaInfo.ContentFilePath = searchPath
			log.Println("mediaInfo.ContentFilePath: ", mediaInfo.ContentFilePath)
			log.Println("--------------------------------------------")
			// log.Println("reconstructManifestFile:", searchPath)
			mediaInfo.ReconstructManifestFile()
			mediaInfo.CreateMediaFileHash()
			genericNonCatFileTree(&mediaInfo, fileTree)

			// if findManifestFile(&mediaInfo) {

			// }
		}

	}
	return nil
}
