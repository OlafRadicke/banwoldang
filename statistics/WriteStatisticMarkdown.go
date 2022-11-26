package statistics

import (
	"os"
	"strconv"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)


// WriteStatisticMarkdown write some statistics in a Markdown file
func (statistic *Statistics) WriteStatisticMarkdown() {
	mdPath := statistic.StatisticDir + "/statistics.md"
	mdFile, err := os.Create(mdPath)

	if err != nil {
		cl.ErrorLogger.Println(err)
		return
	}
	defer mdFile.Close()
	textline := "Count of founded files: " + strconv.Itoa(statistic.FoundedFiles) + "\n"
	textline = textline + "Count of founded categories: " + strconv.Itoa(len(statistic.UsedTags)) + "\n"
	textline = textline + "Count of different name parts: " + strconv.Itoa(len(statistic.PartsOfNames)) + "\n"
	textline = textline + "Count of different Duration: " + strconv.Itoa(len(statistic.Duration)) + "\n"
	textline = textline + "Count of different Height: " + strconv.Itoa(len(statistic.ResolutionHeight)) + "\n"
	textline = textline + "Count of different Width: " + strconv.Itoa(len(statistic.ResolutionWidth)) + "\n"
	_, err = mdFile.WriteString(textline)
	if err != nil {
		cl.ErrorLogger.Println(err)
		return
	}
}
