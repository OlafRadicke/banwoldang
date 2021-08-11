package mediainformation

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"io"
	"os"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// CreateMediaFileHash Create a hash sum of the mediafile for the
// authenticating and to use als target link name.
func (mediaInfo *MediaInformation) CreateMediaFileHash() {

	if mediaInfo.UseChecksum {
		openFile, err := os.Open(mediaInfo.AbsoluteContentSourcePath)
		if err != nil {
			// cl.ErrorLogger.Fatal(err)
			cl.ErrorLogger.Fatal("Can't open file for hashing: ", err)
			defer openFile.Close()
			return
		}
		defer openFile.Close()

		hasher := sha256.New()
		if _, err := io.Copy(hasher, openFile); err != nil {
			cl.ErrorLogger.Fatal(err)
		}

		mediaInfo.SetHashValue(base64.URLEncoding.EncodeToString(hasher.Sum(nil)))

		cl.InfoLogger.Println("REAL HASH: ", mediaInfo.HashValue)

	} else {
		pseutoHash := sha256.Sum256([]byte(mediaInfo.AbsoluteContentSourcePath))
		cl.InfoLogger.Println("PSEUDO HASH: ", hex.EncodeToString(pseutoHash[:]))
		mediaInfo.SetHashValue(hex.EncodeToString(pseutoHash[:]))

		// uuid, err := exec.Command("uuidgen").Output()
		// if err != nil {
		// 	cl.ErrorLogger.Fatal(err)
		// }
		// mediaInfo.HashValue = base64.URLEncoding.EncodeToString(uuid)
		// cl.InfoLogger.Println("UUID base64: ", mediaInfo.HashValue)

	}
	cl.InfoLogger.Println("Hash: ", mediaInfo.HashValue)
}
