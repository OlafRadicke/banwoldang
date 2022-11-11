package mediainformation

import (
	// "encoding/xml"
	// "io/ioutil"
	// "os"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
	gt "github.com/OlafRadicke/go-gthumb"
)

// ReadingManifestFile Read and pars the manifest file
func (mediaInfo *MediaInformation) ReadingManifestFile() {
	var err error
	// xmlFile, err := os.Open(mediaInfo.AbsoluteManifestSourcePath)
	// if err != nil {
	// 	cl.ErrorLogger.Println("Can't read manifest fil: ", err)
	// }

	// byteValue, _ := ioutil.ReadAll(xmlFile)
	// var comment XMLComment
	// if err := xml.Unmarshal(byteValue, &comment); err != nil {
	// 	cl.ErrorLogger.Fatal(err)
	// }

	// for i := 0; i < len(comment.Categories.CategoryList); i++ {
	// 	mediaInfo.Categories = append(mediaInfo.Categories, comment.Categories.CategoryList[i].Value)
	// }
	// defer xmlFile.Close()

	mediaInfo.Comments, err = gt.NewCommentsFile(mediaInfo.AbsoluteManifestSourcePath)
	if err != nil {
		cl.ErrorLogger.Println("error to init object: %w", err)
	}


}
