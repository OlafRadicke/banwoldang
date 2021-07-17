package common

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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

func (fileTree *FileTree) fileHandler(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(path, info.Name(), info.Size())
	if info.IsDir() {
		fmt.Println("Just a directory")
	} else {
		fmt.Println("Is a file! And typ:", filepath.Ext(info.Name()))
		fileTree.Findings++
		if filepath.Ext(info.Name()) == ".xml" {
			fmt.Println("..And a manifet file!!")
		}

	}
	return nil
}
