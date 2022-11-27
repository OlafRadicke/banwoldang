package mediainformation

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
    "fmt"
	"io"
	"os"
	cl "github.com/OlafRadicke/banwoldang/customlogger"
)
func (mediaInfo *MediaInformation) GetHashValue() (string, error){
	var err error
	var hash string
	if mediaInfo.hashValue == "" {
		hash, err = mediaInfo.createMediaFileHash()
		if err != nil {
			return "", err
		}
		mediaInfo.hashValue = hash
	}
	return mediaInfo.hashValue, nil
}

func (mediaInfo *MediaInformation) createMediaFileHash() (string, error) {
	var hashSum string
	var openFile *os.File
	var err error
	if mediaInfo.progConfig.UseChecksum {
		openFile, err = os.Open(mediaInfo.AbsoluteContentSourcePath)
		if err != nil {
			fmt.Println("error! check the log files for more information")
			cl.OrphanLogger.Println(mediaInfo.AbsoluteContentSourcePath)
			cl.ErrorLogger.Fatal("Can't open file for hashing: ", err)
			return "", err
		}
		defer openFile.Close()
		hasher := sha256.New()
		_, err := io.Copy(hasher, openFile)
		if err != nil {
			fmt.Println("error! check the log files for more information")
			cl.ErrorLogger.Fatal(err)
			return "", err
		}
		hashSum = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
		cl.InfoLogger.Println("REAL HASH: ", hashSum)
	} else {
		pseutoHash := sha256.Sum256([]byte(mediaInfo.AbsoluteContentSourcePath))
		cl.InfoLogger.Println("PSEUDO HASH: ", hex.EncodeToString(pseutoHash[:]))
		hashSum = hex.EncodeToString(pseutoHash[:])
	}
	return hashSum, nil
}

