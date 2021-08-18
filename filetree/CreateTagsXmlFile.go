package filetree

import (
	"os"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// fileHandler This function handel and prosessing the file operations.
func (fileTree *FileTree) CreateTagsXmlFile() {
	homePath := os.Getenv("HOME")
	xmlFilePath := homePath + "/.config/gthumb/tags.xml"

	xmlContent := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"
	xmlContent += "<tags version=\"1.0\">\n"
	for key, value := range fileTree.AllUsedCategories {
		cl.InfoLogger.Println("=> ", key, " (", value, ")")
		xmlContent += "\t <tag value=\"" + key + "\"/>\n"
	}
	xmlContent += "</tags>\n"
	cl.InfoLogger.Println(xmlContent)
	openFile, err := os.Create(xmlFilePath)
	if err != nil {
		cl.ErrorLogger.Println(err)
		return
	}
	defer openFile.Close()
	_, err2 := openFile.WriteString(xmlContent)
	if err2 != nil {
		cl.ErrorLogger.Println(err)
		return
	}
	openFile.Sync()
	cl.InfoLogger.Println("rewite file: ", xmlFilePath)
}
