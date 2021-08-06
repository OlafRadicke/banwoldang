package mediainformation

import (
	"crypto/sha256"
	"encoding/base64"
	"io"
	"os"
	"os/exec"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

func (mediaInfo *MediaInformation) CreateMediaFileHash() {
	openFile, err := os.Open(mediaInfo.AbsoluteContentSourcePath)
	if err != nil {
		// cl.ErrorLogger.Fatal(err)
		cl.ErrorLogger.Println("Open file for hashing: ", err)
		defer openFile.Close()
		return
	}
	defer openFile.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, openFile); err != nil {
		cl.ErrorLogger.Fatal(err)
	}
	ifTest := false
	if ifTest {
		uuid, err := exec.Command("uuidgen").Output()
		if err != nil {
			cl.ErrorLogger.Fatal(err)
		}
		mediaInfo.HashValue = base64.URLEncoding.EncodeToString(uuid)
		cl.InfoLogger.Println("UUID base64: ", mediaInfo.HashValue)
	} else {
		mediaInfo.HashValue = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
		cl.InfoLogger.Println("hash: ", mediaInfo.HashValue)
	}
	// cl.InfoLogger.Println("Hash: ", mediaInfo.HashValue)
}
