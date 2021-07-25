package mediainformation

import (
	"fmt"
	"log"
	"os"
)

// Create an empty manifest file
func (mediaInfo *MediaInformation) CreateEmptyManifestFile() {
	fmt.Println("CreateEmptyManifestFile()")
	f, err := os.Create(mediaInfo.AbsoluteManifestSourcePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	rawtext := `<?xml version="1.0" encoding="UTF-8"?>
<comment version="3.0">
  <caption/>
  <note/>
  <place/>
  <categories>
    <category value="00-script-create-manifest"/>
  </categories>
</comment>`

	_, err2 := f.WriteString(rawtext)

	if err2 != nil {
		log.Fatal(err2)
	}

}
