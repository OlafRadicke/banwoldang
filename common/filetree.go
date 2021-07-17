package common

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type filetree interface {
	Collecting()
	PrintAll()
	fileHandler(path string, info os.FileInfo, err error) error
}

type FileTree struct {
	StartPath string
	Findings  int
}

func (fileTree *FileTree) Collecting() {
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
		if filepath.Ext(info.Name()) == ".xml" {
			log.Println("===========================================")
			log.Println("Manifest file: ", searchPath)
			log.Println("Content file: ", reconstructContenFile(searchPath))
		}

	}
	return nil
}

//  Reconstruct the path of the conten file from the path of a manifest file.
func reconstructContenFile(manifestFilePath string) string {
	pathParts := strings.Split(manifestFilePath, "/")
	onsUpLevel := len(pathParts) - 2
	onsUpPath := strings.Join(pathParts[0:onsUpLevel], "/")

	manifestCoreName := pathParts[len(pathParts)-1]
	extension := filepath.Ext(manifestCoreName)
	contentFileName := manifestCoreName[0 : len(manifestCoreName)-len(extension)]

	contenFilePath := onsUpPath + "/" + contentFileName
	if _, err := os.Stat(contenFilePath); err == nil {
		// path/to/whatever exists
		log.Println("Manifet file and conten file a exit!")
	} else {
		log.Println("File not exist: ", contenFilePath)
	}
	return contenFilePath
}
