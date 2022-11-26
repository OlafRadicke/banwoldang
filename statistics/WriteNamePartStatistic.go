package statistics

import (
	"os"
	"strconv"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)


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
