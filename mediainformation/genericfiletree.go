package mediainformation

import (
	"log"
	"os"
	"path/filepath"
)

// Generate directory tree with symlinks of file with categories.
func (mediaInfo *MediaInformation) GenericFileTree(genericDir string) {

	for i := 0; i < len(mediaInfo.Categories); i++ {
		mediaInfo.GenerateAbsoluteLinkDirContentPath(genericDir, mediaInfo.Categories[i])

		absolutLinkSource, err2 := filepath.Abs(mediaInfo.ContentFilePath)
		if err2 != nil {
			log.Fatal(err2)
		}
		// log.Println("mediaInfo.ContentFilePath: ", absolutLinkSource)

		katPath := genericDir + "gereric-tree/" + mediaInfo.Categories[i]
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

}
