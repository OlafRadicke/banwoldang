package mediainformation

import (
	"log"
	"os"
)

// Generate directory with symlinks of file without categories.
func (mediaInfo *MediaInformation) GenericNonCatFileTree(genericDir string) {

	mediaInfo.SetAbsoluteContentLinkDirPath(genericDir, "00-no-cats")

	katPath := genericDir + "gereric-tree/00-no-cats/"
	if _, err := os.Stat(katPath); os.IsNotExist(err) {
		err := os.MkdirAll(katPath, 0770)
		if err != nil {
			log.Fatal(err)
		}
	}

	err := os.Symlink(mediaInfo.AbsoluteContentSourcePath, mediaInfo.AbsoluteContentLinkDirPath)
	if err != nil {
		// log.Fatal("Create symlink: ", err)
		log.Println("Create symlink: ", err)
		return
	}
}
