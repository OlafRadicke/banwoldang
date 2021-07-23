package common

import (
	"log"
	"os"
	"path/filepath"

	"github.com/OlafRadicke/banwoldang/mediainformation"
)

// Generate directory with symlinks of file without categories.
func genericNonCatFileTree(mediaInfo *mediainformation.MediaInformation, fileTree *FileTree) {
	log.Println("Generate directory with file without categories. ")
	genFilePath := fileTree.GenericDir + "gereric-tree/00-no-cats/" + mediaInfo.HashValue + mediaInfo.Extension
	absolutLinkTarget, err1 := filepath.Abs(genFilePath)
	if err1 != nil {
		log.Fatal(err1)
	}

	absolutLinkSource, err2 := filepath.Abs(mediaInfo.ContentFilePath)
	if err2 != nil {
		log.Fatal(err2)
	}
	// log.Println("mediaInfo.ContentFilePath: ", absolutLinkSource)

	katPath := fileTree.GenericDir + "gereric-tree/00-no-cats/"
	if _, err := os.Stat(katPath); os.IsNotExist(err) {
		err := os.MkdirAll(katPath, 0770)
		if err != nil {
			log.Fatal(err)
		}
	}

	err := os.Symlink(absolutLinkSource, absolutLinkTarget)
	if err != nil {
		// log.Fatal("Create symlink: ", err)
		log.Println("Create symlink: ", err)
		return
	}
}

func genericSingleCatFileTree(mediaInfo *mediainformation.MediaInformation, fileTree *FileTree) {
	// pass
}

// Generate directory tree with symlinks of file with categories.
func genericFileTree(mediaInfo *mediainformation.MediaInformation, fileTree *FileTree) {

	for i := 0; i < len(mediaInfo.Categories); i++ {

		genFilePath := fileTree.GenericDir + "gereric-tree/" + mediaInfo.Categories[i] + "/" + mediaInfo.HashValue + mediaInfo.Extension
		absolutLinkTarget, err1 := filepath.Abs(genFilePath)
		if err1 != nil {
			log.Fatal(err1)
		}
		// log.Println("genFilePath: ", absolutLinkTarget)

		absolutLinkSource, err2 := filepath.Abs(mediaInfo.ContentFilePath)
		if err2 != nil {
			log.Fatal(err2)
		}
		// log.Println("mediaInfo.ContentFilePath: ", absolutLinkSource)

		katPath := fileTree.GenericDir + "gereric-tree/" + mediaInfo.Categories[i]
		if _, err := os.Stat(katPath); os.IsNotExist(err) {
			err := os.MkdirAll(katPath, 0770)
			if err != nil {
				log.Fatal(err)
			}
		}

		err := os.Symlink(absolutLinkSource, absolutLinkTarget)
		if err != nil {
			// log.Fatal("Create symlink: ", err)
			log.Println("Create symlink: ", err)
			return
		}
	}

}
