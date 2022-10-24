package statistics

import (
	"log"
	"os"
	"strconv"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// statistics The inteface of the feletree struckt
type statistics interface {
	GoThroughCollection()
	ContPartsOfNames(string)
	writeStatistic(string)
}

type Statistics struct {
	UsedTags         map[string]int
	PartsOfNames     map[string]int
	Duration         map[int]int
	ResolutionHeight map[int]int
	ResolutionWidth  map[int]int
	FoundedFiles     int
}

// NewStatistics create new instance of Statistics and get it back.
func NewStatistics() *Statistics {
	statistic := Statistics{}
	statistic.FoundedFiles = 0
	statistic.UsedTags = make(map[string]int)
	statistic.PartsOfNames = make(map[string]int)
	statistic.Duration = make(map[int]int)
	statistic.ResolutionHeight = make(map[int]int)
	statistic.ResolutionWidth = make(map[int]int)
	return &statistic
}

// ContPartsOfNames count part of file names.
func (statistic *Statistics) ContPartsOfNames(namePart string) {
	if count, ok := statistic.PartsOfNames[namePart]; ok {
		cl.InfoLogger.Println("The name part ", namePart, " is allready added (", count, "). Count up...")
		statistic.PartsOfNames[namePart]++
	} else {
		cl.InfoLogger.Println("The name part ", namePart, " (", count, ") is new.")
		statistic.PartsOfNames[namePart] = 1
	}
}

// WriteStatistic write files with statistics
func (statistic *Statistics) WriteStatistic(statisticDir string) {
	var key string
	var value int

	_, err := os.Stat(statisticDir)
	if os.IsNotExist(err) {
		err := os.MkdirAll(statisticDir, 0770)
		if err != nil {
			cl.ErrorLogger.Fatal(err)
		}
	}

	// ------- CVS BEGIN -----------
	csvPath := statisticDir + "./statistics.csv"
	csvFile, err := os.Create(csvPath)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	for key, value = range statistic.UsedTags {
		textline := "\"" + key + "\";\"" + strconv.Itoa(value) + "\"\n"
		_, err = csvFile.WriteString(textline)
		if err != nil {
			cl.ErrorLogger.Println(err)
			return
		}
	}
	// ------- CVS END -----------

	// ------- MARKDOWN BEGIN -----------
	mdPath := statisticDir + "./statistics.md"
	mdFile, err := os.Create(mdPath)
	if err != nil {
		log.Fatal(err)
	}
	defer mdFile.Close()
	textline := "Count of founded files: " + strconv.Itoa(statistic.FoundedFiles) + "\n"
	textline = textline + "Count of founded categories: " + strconv.Itoa(len(statistic.UsedTags)) + "\n"
	_, err = mdFile.WriteString(textline)
	if err != nil {
		cl.ErrorLogger.Println(err)
		return
	}
	// ------- MARKDOWN ENDE -----------

}
