package common

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// XML strctur of manifest file:
//
// <?xml version="1.0" encoding="UTF-8"?>
// <comment version="3.0">
//   <caption/>
//   <note/>
//   <place/>
//   <categories>
//     <category value="c-anal"/>
//     <category value="f-athletic"/>
//     <category value="f-brunette"/>
//     <category value="l-living_room"/>
//     <category value="n-german"/>
//     <category value="q-90s"/>
//     <category value="q-aaa"/>
//     <category value="q-aaaa"/>
//     <category value="q-self_made"/>
//     <category value="q-video_hd"/>
//     <category value="s-bottomless"/>
//     <category value="s-chav"/>
//     <category value="s-socks"/>
//     <category value="s-tattoo"/>
//   </categories>
// </comment>

type Comment struct {
	XMLName    xml.Name   `xml:"comment"`
	Caption    string     `xml:"caption"`
	Note       string     `xml:"note"`
	Place      string     `xml:"place"`
	Categories Categories `xml:"categories"`
}

type Categories struct {
	CategoryList []Category `xml:"category"`
	Value        string     `xml:"value,attr"`
}

type Category struct {
	Value string `xml:"value,attr"`
}

func readingManifestFile(mediaInfo *MediaInformation) {
	xmlFile, err := os.Open(mediaInfo.ManifestFilePath)
	if err != nil {
		log.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var comment Comment
	if err := xml.Unmarshal(byteValue, &comment); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(comment.Categories.CategoryList); i++ {
		// fmt.Println("cat value: " + comment.Categories.CategoryList[i].Value)
		mediaInfo.Categories = append(mediaInfo.Categories, comment.Categories.CategoryList[i].Value)
	}

	fmt.Println("cat len: ", len(mediaInfo.Categories))
	defer xmlFile.Close()

}

// func findManifestFile(mediaInfo *MediaInformation) bool {
// 	mediaInfo.ContentFilePath

// 	testManifestFilePath = mediaInfo.ContentFilePath + ".comments/" + filename + ".xml"
// }

//  Reconstruct the path of the conten file from the path of a manifest file.
func reconstructManifestFile(mediaInfo *MediaInformation) {
	pathParts := strings.Split(mediaInfo.ContentFilePath, "/")
	mediaInfo.ContentFileName = pathParts[len(pathParts)-1]
	baseDir := filepath.Dir(mediaInfo.ContentFilePath)
	mediaInfo.Extension = filepath.Ext(mediaInfo.ContentFileName)
	mediaInfo.ManifestFilePath = baseDir + ".comments/" + mediaInfo.ContentFileName + ".xml"

	if _, err := os.Stat(mediaInfo.ContentFilePath); err == nil {
		// path/to/whatever exists
		// log.Println("Manifet file and conten file a exit!")
	} else {
		log.Println("Manifest file not exist: ", mediaInfo.ManifestFilePath)
	}

}
