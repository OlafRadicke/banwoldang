package statistics

import (
	"os"
	"strconv"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// WriteDurationStatistic creates an csv file withe the statistic of Duration
func (statistic *Statistics) WriteDurationStatistic() {
	var key int
	var value int
	var textline string

	durationPath := statistic.StatisticDir + "/statistics_duration.csv"
	durationFile, err := os.Create(durationPath)
	if err != nil {
		cl.ErrorLogger.Println(err)
		return
	}
	defer durationFile.Close()

	textline = "\"minutes\";\"count\"\n"
	_, err = durationFile.WriteString(textline)
	if err != nil {
		cl.ErrorLogger.Println(err)
		return
	}

	for key, value = range statistic.Duration {
		textline = "\"" + strconv.Itoa(key) + "\";\"" + strconv.Itoa(value) + "\"\n"
		_, err = durationFile.WriteString(textline)
		if err != nil {
			cl.ErrorLogger.Println(err)
			return
		}
	}

}

// WriteResolutionHeightStatistic creates an csv file withe the statistic
// about Height of the Resolution
func (statistic *Statistics) WriteResolutionHeightStatistic() {
	var key int
	var value int
	var textline string

	resolutionPath := statistic.StatisticDir + "/statistics_resolution_height.csv"
	resolutionFile, err := os.Create(resolutionPath)
	if err != nil {
		cl.ErrorLogger.Println(err)
		return
	}
	defer resolutionFile.Close()

	textline = "\"pixel height\";\"count\"\n"
	_, err = resolutionFile.WriteString(textline)
	if err != nil {
		cl.ErrorLogger.Println(err)
		return
	}

	for key, value = range statistic.ResolutionHeight {
		textline = "\"" + strconv.Itoa(key) + "\";\"" + strconv.Itoa(value) + "\"\n"
		_, err = resolutionFile.WriteString(textline)
		if err != nil {
			cl.ErrorLogger.Println(err)
			return
		}
	}

}

// WriteResolutionWidthStatistic creates an csv file withe the statistic
// about Width of the Resolution
func (statistic *Statistics) WriteResolutionWidthStatistic() {
	var key int
	var value int
	var textline string

	resolutionPath := statistic.StatisticDir + "/statistics_resolution_width.csv"
	resolutionFile, err := os.Create(resolutionPath)
	if err != nil {
		cl.ErrorLogger.Println(err)
		return
	}
	defer resolutionFile.Close()

	textline = "\"pixel width\";\"count\"\n"
	_, err = resolutionFile.WriteString(textline)
	if err != nil {
		cl.ErrorLogger.Println(err)
		return
	}
	for key, value = range statistic.ResolutionWidth {
		textline = "\"" + strconv.Itoa(key) + "\";\"" + strconv.Itoa(value) + "\"\n"
		_, err = resolutionFile.WriteString(textline)
		if err != nil {
			cl.ErrorLogger.Println(err)
			return
		}
	}

}

// WriteTagStatistic creates an csv file withe the statistic of the used tags
func (statistic *Statistics) WriteTagStatistic() {
	var key string
	var value int
	var textline string

	csvPath := statistic.StatisticDir + "/statistics_tag.csv"
	csvFile, err := os.Create(csvPath)
	if err != nil {
		cl.ErrorLogger.Println(err)
		return
	}
	defer csvFile.Close()

	textline = "\"tag\";\"count\"\n"
	_, err = csvFile.WriteString(textline)
	if err != nil {
		cl.ErrorLogger.Println(err)
		return
	}
	for key, value = range statistic.UsedTags {
		textline = "\"" + key + "\";\"" + strconv.Itoa(value) + "\"\n"
		_, err = csvFile.WriteString(textline)
		if err != nil {
			cl.ErrorLogger.Println(err)
			return
		}
	}

}

// WriteNamePartStatistic creates an csv file withe the statistic of
// the parts of the file names
func (statistic *Statistics) WriteNamePartStatistic() {
	csvPartsPath := statistic.StatisticDir + "/statistics_name-parts.csv"
	csvPartsFile, err := os.Create(csvPartsPath)
	var key string
	var value int
	var textline string

	if err != nil {
		cl.ErrorLogger.Println(err)
		return
	}
	defer csvPartsFile.Close()

	textline = "\"name part\";\"count\"\n"
	_, err = csvPartsFile.WriteString(textline)
	if err != nil {
		cl.ErrorLogger.Println(err)
		return
	}
	for key, value = range statistic.PartsOfNames {
		textline = "\"" + key + "\";\"" + strconv.Itoa(value) + "\"\n"
		_, err = csvPartsFile.WriteString(textline)
		if err != nil {
			cl.ErrorLogger.Println(err)
			return
		}
	}
}

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

// WriteStatistic write files with statistics
func (statistic *Statistics) WriteStatistic() {

	_, err := os.Stat(statistic.StatisticDir)
	if os.IsNotExist(err) {
		err := os.MkdirAll(statistic.StatisticDir, 0770)
		if err != nil {
			return
			cl.ErrorLogger.Fatal(err)
		}
	}
	statistic.WriteDurationStatistic()
	statistic.WriteNamePartStatistic()
	statistic.WriteResolutionHeightStatistic()
	statistic.WriteResolutionWidthStatistic()
	statistic.WriteStatisticMarkdown()
	statistic.WriteTagStatistic()

}
