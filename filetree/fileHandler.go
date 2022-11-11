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
		fileTree.Statistic.FoundedFiles++

		if filepath.Ext(info.Name()) == ".xml" {
			mediaInfo := mediainformation.NewMediaInformationByManifest(fileTree.progConfig, fileTree.Statistic, searchPath)
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
			mediaInfo := mediainformation.NewMediaInformation(fileTree.progConfig, fileTree.Statistic, searchPath)
			mediaInfo.SetAbsoluteContentSourcePath(searchPath)
			mediaInfo.ReconstructManifestFile()
			// mediaInfo.CreateMediaFileHash()
			mediaInfo.GenerateLinkDirTreeWithoutManifests()

		}
		return nil
	}
}
