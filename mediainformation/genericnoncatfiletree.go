package mediainformation

import (
	"log"
	"os"
	"path/filepath"
)

// Generate directory with symlinks of file without categories.
func (mediaInfo *MediaInformation) GenericNonCatFileTree(genericDir string) {

	mediaInfo.GenerateAbsoluteLinkDirContentPath(genericDir, "00-no-cats")

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

	err := os.Symlink(absolutLinkSource, mediaInfo.AbsoluteLinkDirContentTarget)
	if err != nil {
		// log.Fatal("Create symlink: ", err)
		log.Println("Create symlink: ", err)
		return
	}
}
