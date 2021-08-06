package mediainformation

import (
	"encoding/xml"
	"io/ioutil"
	"os"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

func (mediaInfo *MediaInformation) ReadingManifestFile() {
	xmlFile, err := os.Open(mediaInfo.AbsoluteManifestSourcePath)
	if err != nil {
		cl.InfoLogger.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var comment XMLComment
	if err := xml.Unmarshal(byteValue, &comment); err != nil {
		cl.ErrorLogger.Fatal(err)
	}

	for i := 0; i < len(comment.Categories.CategoryList); i++ {
		// fmt.Println("cat value: " + comment.Categories.CategoryList[i].Value)
		mediaInfo.Categories = append(mediaInfo.Categories, comment.Categories.CategoryList[i].Value)
	}
	defer xmlFile.Close()
}
