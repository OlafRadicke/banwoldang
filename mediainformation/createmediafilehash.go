package mediainformation

import (
	"crypto/sha256"
	"encoding/base64"
	"io"
	"log"
	"os"
	"os/exec"
)

func (mediaInfo *MediaInformation) CreateMediaFileHash() {
	openFile, err := os.Open(mediaInfo.AbsoluteContentSourcePath)
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
	ifTest := false
	if ifTest {
		uuid, err := exec.Command("uuidgen").Output()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("UUID: ", uuid)
		mediaInfo.HashValue = base64.URLEncoding.EncodeToString(uuid)
		log.Println("UUID base64: ", mediaInfo.HashValue)
	} else {
		mediaInfo.HashValue = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
		log.Println("hash: ", mediaInfo.HashValue)
	}
	// log.Println("Hash: ", mediaInfo.HashValue)
}
