package mediainformation

import (
	"os"
	"path/filepath"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
	gt "github.com/OlafRadicke/go-gthumb"
)

// CreateNewEmptyManifestFile Create an new empty manifest file it is not exist
// TODO use Lib github.com/OlafRadicke/go-gthumb
func (mediaInfo *MediaInformation) CreateNewEmptyManifestFile() {

	_, err := os.Stat(mediaInfo.AbsoluteManifestSourcePath)
	if os.IsNotExist(err) {
		cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		cl.InfoLogger.Println("+++++++++++++++++++ CREAT MANIFEST +++++++++++++++++++++")
		cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		cl.InfoLogger.Println(err, mediaInfo.AbsoluteManifestSourcePath)

		manifestBaseDir := filepath.Dir(mediaInfo.AbsoluteManifestSourcePath)
		err := os.MkdirAll(manifestBaseDir, 0770)
		if err != nil {
			cl.ErrorLogger.Fatal("[202108040831]", err)
		}

		comment, err := gt.NewCommentsFile(mediaInfo.AbsoluteManifestSourcePath)
		if err != nil {
			cl.ErrorLogger.Println("error to open comment file: ", err)
			cl.ErrorLogger.Println("try later to create an new: ", mediaInfo.AbsoluteManifestSourcePath)
		}

		comment.AddCategory("00-script-create-manifest")

		err = comment.Save()
		if err != nil {
			cl.ErrorLogger.Fatal("error writting comment file: %w", err)
		}

// 		openFile, err := os.Create(mediaInfo.AbsoluteManifestSourcePath)
// 		if err != nil {
// 			cl.ErrorLogger.Fatal("[202108040825]", err)
// 		}
// 		defer openFile.Close()
// 		rawtext := `<?xml version="1.0" encoding="UTF-8"?>
// <comment version="3.0">
// 	<caption/>
// 	<note/>
// 	<place/>
// 	<categories>
// 	<category value="00-script-create-manifest"/>
// 	</categories>
// </comment>`
// 		cl.InfoLogger.Println(rawtext)
// 		_, err2 := openFile.WriteString(rawtext)

// 		if err2 != nil {
// 			cl.ErrorLogger.Fatal("[202108040824]", err2)
// 		}
	}

}
