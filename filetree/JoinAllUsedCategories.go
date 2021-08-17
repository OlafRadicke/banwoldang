package filetree

import (
	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// fileHandler This function handel and prosessing the file operations.
func (fileTree *FileTree) JoinAllUsedCategories(categories []string) {

	for _, cat := range categories {
		if count, ok := fileTree.AllUsedCategories[cat]; ok {
			cl.InfoLogger.Println("The category ", cat, " is allready added (", count, "). Count up...")
			fileTree.AllUsedCategories[cat]++
		} else {
			cl.InfoLogger.Println("The category ", cat, " (", count, ") is new.")
			fileTree.AllUsedCategories[cat] = 1
		}
	}

}
