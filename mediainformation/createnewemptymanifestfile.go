package mediainformation

import (
	"log"
	"os"
	"path/filepath"
)

// Create an new empty manifest file it is not exist
func (mediaInfo *MediaInformation) CreateNewEmptyManifestFile() {

	_, err := os.Stat(mediaInfo.AbsoluteManifestSourcePath)
	if os.IsNotExist(err) {
		log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		log.Println("+++++++++++++++++++ CREAT MANIFEST +++++++++++++++++++++")
		log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		log.Println(err, mediaInfo.AbsoluteManifestSourcePath)

		manifestBaseDir := filepath.Dir(mediaInfo.AbsoluteManifestSourcePath)
		err := os.MkdirAll(manifestBaseDir, 0770)
		if err != nil {
			log.Fatal("[202108040831]", err)
		}

		openFile, err := os.Create(mediaInfo.AbsoluteManifestSourcePath)
		if err != nil {
			log.Fatal("[202108040825]", err)
		}
		defer openFile.Close()
		rawtext := `<?xml version="1.0" encoding="UTF-8"?>
<comment version="3.0">
	<caption/>
	<note/>
	<place/>
	<categories>
	<category value="00-script-create-manifest"/>
	</categories>
</comment>`
		log.Println(rawtext)
		_, err2 := openFile.WriteString(rawtext)

		if err2 != nil {
			log.Fatal("[202108040824]", err2)
		}
	}

}
