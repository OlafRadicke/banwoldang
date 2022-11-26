package statistics

import (
	"os"
	"strconv"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)


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
