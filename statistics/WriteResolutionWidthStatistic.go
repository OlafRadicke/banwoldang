package statistics

import (
	"os"
	"strconv"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)



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
