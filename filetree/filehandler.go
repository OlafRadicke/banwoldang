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
			absolutSearchPath, err2 := filepath.Abs(searchPath)
			if err2 != nil {
				log.Fatal(err2)
			}
			mediaInfo.AbsoluteManifestSourcePath = absolutSearchPath
			log.Println("Manifest file: ", searchPath)
			log.Println("mediaInfo.AbsoluteManifestSourcePath: ", mediaInfo.AbsoluteManifestSourcePath)
			log.Println("===========================================")

			mediaInfo.SetAbsoluteLinkDirPath(fileTree.LinkDir)
			mediaInfo.ReconstructContenSourceFile()
			mediaInfo.ReadingManifestFile()
			mediaInfo.CreateMediaFileHash()
			mediaInfo.GenerateFileTree()
		} else {
			log.Println("-------------------------------------------")
			log.Println("searchPath: ", searchPath)
			contentSourcePath, err2 := filepath.Abs(searchPath)
			if err2 != nil {
				log.Fatal(err2)
			}

			log.Println("contentSourcePath: ", contentSourcePath)
			mediaInfo.AbsoluteContentSourcePath = contentSourcePath

			log.Println("--------------------------------------------")
			// log.Println("reconstructManifestFile:", searchPath)

			mediaInfo.SetAbsoluteLinkDirPath(fileTree.LinkDir)
			mediaInfo.ReconstructManifestFile()
			mediaInfo.CreateMediaFileHash()
			mediaInfo.GenerateNonCatFileTree()

			// if findManifestFile(&mediaInfo) {

			// }
		}

	}
	return nil
}
