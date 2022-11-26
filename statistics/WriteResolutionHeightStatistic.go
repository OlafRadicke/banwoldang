package statistics

import (
	"os"
	"strconv"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)



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
