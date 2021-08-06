package filetree

import (
	"os"
	"path/filepath"

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
			mediaInfo.ReconstructContenSourceFile()
			mediaInfo.ReadingManifestFile()
			mediaInfo.CreateMediaFileHash()
			mediaInfo.GenerateLinkDirTree()
		} else {
			mediaInfo.SetAbsoluteContentSourcePath(searchPath)
			mediaInfo.ReconstructManifestFile()
			mediaInfo.CreateMediaFileHash()
			mediaInfo.GenerateLinkDirTreeWithoutManifests()
		}
		return nil
	}
}
