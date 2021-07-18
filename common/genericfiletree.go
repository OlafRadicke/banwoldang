package common

import (
	"crypto/sha256"
	"encoding/base64"
	"io"
	"log"
	"os"
	"path/filepath"
)

func genericFileTree(mediaInfo *MediaInformation) {

	for i := 0; i < len(mediaInfo.Categories); i++ {
		genFilePath := "./gereric-tree/" + mediaInfo.Categories[i] + "/" + mediaInfo.HashValue + mediaInfo.Extension
		katPath := "./gereric-tree/" + mediaInfo.Categories[i]
		absolutTarget, err1 := filepath.Abs(genFilePath)
		if err1 != nil {
			log.Fatal(err1)
		}

		log.Println("genFilePath: ", absolutTarget)
		absolutSource, err2 := filepath.Abs(mediaInfo.ContentFilePath)
		if err2 != nil {
			log.Fatal(err2)
		}
		log.Println("mediaInfo.ContentFilePath: ", absolutSource)

		if _, err := os.Stat(katPath); os.IsNotExist(err) {
			err := os.MkdirAll(katPath, 0770)
			if err != nil {
				log.Fatal(err)
			}
		}
		err := os.Symlink(absolutSource, absolutTarget)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func createMediaFileHash(mediaInfo *MediaInformation) {
	openFile, err := os.Open(mediaInfo.ContentFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer openFile.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, openFile); err != nil {
		log.Fatal(err)
	}
	mediaInfo.HashValue = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	// log.Println("Hash: ", mediaInfo.HashValue)
}
