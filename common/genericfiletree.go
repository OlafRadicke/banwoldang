package common

import (
	"crypto/sha256"
	"encoding/base64"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Generate directory with symlinks of file without categories.
func genericNonCatFileTree(mediaInfo *MediaInformation, fileTree *FileTree) {
	log.Println("Generate directory with file without categories. ")
}

// Generate directory tree with symlinks of file with categories.
func genericFileTree(mediaInfo *MediaInformation, fileTree *FileTree) {

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

func createMediaFileHash(mediaInfo *MediaInformation) {
	openFile, err := os.Open(mediaInfo.ContentFilePath)
	if err != nil {
		// log.Fatal(err)
		log.Println("Open file for hashing: ", err)
		defer openFile.Close()
		return
	}
	defer openFile.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, openFile); err != nil {
		log.Fatal(err)
	}
	mediaInfo.HashValue = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	// log.Println("Hash: ", mediaInfo.HashValue)
}
