package mediainformation

import (
	"log"
	"os"
	"path/filepath"
)

// Generate directory with symlinks of file without categories.
func (mediaInfo *MediaInformation) GenericNonCatFileTree(genericDir string) {
	log.Println("Generate directory with file without categories. ")
	genFilePath := genericDir + "gereric-tree/00-no-cats/" + mediaInfo.HashValue + mediaInfo.Extension
	absolutLinkTarget, err1 := filepath.Abs(genFilePath)
	if err1 != nil {
		log.Fatal(err1)
	}

	absolutLinkSource, err2 := filepath.Abs(mediaInfo.ContentFilePath)
	if err2 != nil {
		log.Fatal(err2)
	}
	// log.Println("mediaInfo.ContentFilePath: ", absolutLinkSource)

	katPath := genericDir + "gereric-tree/00-no-cats/"
	if _, err := os.Stat(katPath); os.IsNotExist(err) {
		err := os.MkdirAll(katPath, 0770)
		if err != nil {
			log.Fatal(err)
		}
	}

	err := os.Symlink(absolutLinkSource, absolutLinkTarget)
	if err != nil {
		// log.Fatal("Create symlink: ", err)
		log.Println("Create symlink: ", err)
		return
	}
}
