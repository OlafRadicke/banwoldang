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
