package filetree

import (
	cl "github.com/OlafRadicke/banwoldang/customlogger"
	gt "github.com/OlafRadicke/go-gthumb"

)

// fileHandler This function handel and prosessing the file operations.
func (fileTree *FileTree) JoinAllUsedCategories(categories []gt.XmlCategory) {

	for _, cat := range categories {
		if count, ok := fileTree.Statistic.UsedTags[cat.Value]; ok {
			cl.InfoLogger.Println("The category ", cat.Value, " is allready added (", count, "). Count up...")
			fileTree.Statistic.UsedTags[cat.Value]++
		} else {
			cl.InfoLogger.Println("The category ", cat.Value, " (", count, ") is new.")
			fileTree.Statistic.UsedTags[cat.Value] = 1
		}
	}

}
