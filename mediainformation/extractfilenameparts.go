package mediainformation

import (
	"log"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func (mediaInfo *MediaInformation) ExtractFileNameParts() []string {
	var listOfParts []string
	var finaListOfParts []string
	contenExtension := filepath.Ext(mediaInfo.ContentFileName)
	contentCoreName := mediaInfo.ContentFileName[0 : len(mediaInfo.ContentFileName)-len(contenExtension)]

	reg := regexp.MustCompile(`[^a-zA-Z0-9/_]`)
	cleanString := reg.ReplaceAllString(contentCoreName, "_")

	log.Println("cleanString: ", cleanString)
	listOfParts = strings.Split(cleanString, "_")
	for i := 0; i < len(listOfParts); i++ {
		if len(listOfParts[i]) < 4 {
			continue
		}
		_, err := strconv.Atoi(listOfParts[i])
		if err == nil {
			continue
		}

		finaListOfParts = append(finaListOfParts, listOfParts[i])
	}
	return finaListOfParts
}
