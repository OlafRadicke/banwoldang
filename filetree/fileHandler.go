package filetree

import (
	"os"
	"path/filepath"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
	"github.com/OlafRadicke/banwoldang/mediainformation"
)

// fileHandler This function handel and prosessing the file operations.
func (fileTree *FileTree) fileHandler(searchPath string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		// cl.InfoLogger.Println("Just a directory")
		return nil
	} else {
		fileTree.Findings++
		mediaInfo := mediainformation.MediaInformation{}
		mediaInfo.UseChecksum = fileTree.UseChecksum
		mediaInfo.SetAbsoluteLinkDirPath(fileTree.LinkDir)

		if filepath.Ext(info.Name()) == ".xml" {
			mediaInfo.SetAbsoluteManifestSourcePath(searchPath)
			err := mediaInfo.ReconstructContenSourceFile()
			if err != nil {
				cl.ErrorLogger.Println("This manifest file has no media file: ", mediaInfo.AbsoluteManifestSourcePath)
				return nil
			}
			mediaInfo.ReadingManifestFile()
			mediaInfo.CreateMediaFileHash()
			mediaInfo.GenerateLinkDirTree()
		} else {
			mediaInfo.SetAbsoluteContentSourcePath(searchPath)
			mediaInfo.ReconstructManifestFile()
			mediaInfo.CreateMediaFileHash()
			mediaInfo.GenerateLinkDirTreeWithoutManifests()
		}
		for _, cat := range mediaInfo.Categories {
			if count, ok := fileTree.AllUsedCategories[cat]; ok {
				cl.InfoLogger.Println("The category ", cat, " is allready added (", count, "). Count up...")
				fileTree.AllUsedCategories[cat]++
			} else {
				cl.InfoLogger.Println("The category ", cat, " (", count, ") is new.")
				fileTree.AllUsedCategories[cat] = 1
			}
		}
		return nil
	}
}
