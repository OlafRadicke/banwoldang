package filetree

import (
	"log"
	"os"
	"path/filepath"

	"github.com/OlafRadicke/banwoldang/mediainformation"
)

func (fileTree *FileTree) fileHandler(searchPath string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	log.Println("func params:", searchPath, info.Name(), info.Size())
	if info.IsDir() {
		// log.Println("Just a directory")
	} else {
		fileTree.Findings++
		mediaInfo := mediainformation.MediaInformation{}
		if filepath.Ext(info.Name()) == ".xml" {

			log.Println("===========================================")
			log.Println("searchPath: ", searchPath)
			mediaInfo.ManifestFilePath = searchPath
			log.Println("Manifest file: ", searchPath)
			log.Println("mediaInfo.ManifestFilePath: ", mediaInfo.ManifestFilePath)
			log.Println("===========================================")

			mediaInfo.ReconstructContenFile()
			mediaInfo.ReadingManifestFile()
			mediaInfo.CreateMediaFileHash()
			mediaInfo.GenericFileTree(fileTree.GenericDir)
		} else {
			log.Println("-------------------------------------------")
			log.Println("searchPath: ", searchPath)
			mediaInfo.ContentFilePath = searchPath
			log.Println("mediaInfo.ContentFilePath: ", mediaInfo.ContentFilePath)
			log.Println("--------------------------------------------")
			// log.Println("reconstructManifestFile:", searchPath)
			mediaInfo.ReconstructManifestFile()
			mediaInfo.CreateMediaFileHash()
			mediaInfo.GenericNonCatFileTree(fileTree.GenericDir)

			// if findManifestFile(&mediaInfo) {

			// }
		}

	}
	return nil
}
