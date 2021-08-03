package mediainformation

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

func (mediaInfo *MediaInformation) ReadingManifestFile() {
	xmlFile, err := os.Open(mediaInfo.AbsoluteManifestSourcePath)
	if err != nil {
		log.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var comment XMLComment
	if err := xml.Unmarshal(byteValue, &comment); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(comment.Categories.CategoryList); i++ {
		// fmt.Println("cat value: " + comment.Categories.CategoryList[i].Value)
		mediaInfo.Categories = append(mediaInfo.Categories, comment.Categories.CategoryList[i].Value)
	}
	defer xmlFile.Close()
}
