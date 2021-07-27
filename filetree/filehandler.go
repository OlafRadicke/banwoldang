package filetree

import (
	"os"
	"path/filepath"

	"github.com/OlafRadicke/banwoldang/mediainformation"
)

func (fileTree *FileTree) fileHandler(searchPath string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		// log.Println("Just a directory")
		return nil
	} else {
		fileTree.Findings++
		mediaInfo := mediainformation.MediaInformation{}
		mediaInfo.SetAbsoluteLinkDirPath(fileTree.LinkDir)

		if filepath.Ext(info.Name()) == ".xml" {
			mediaInfo.SetAbsoluteManifestSourcePath(searchPath)
			mediaInfo.ReconstructContenSourceFile()
			mediaInfo.ReadingManifestFile()
			mediaInfo.CreateMediaFileHash()
			mediaInfo.GenerateFileTree()
		} else {
			mediaInfo.SetAbsoluteContentSourcePath(searchPath)
			mediaInfo.ReconstructManifestFile()
			mediaInfo.CreateMediaFileHash()
			mediaInfo.GenerateNonCatFileTree()
		}
		return nil
	}
}
