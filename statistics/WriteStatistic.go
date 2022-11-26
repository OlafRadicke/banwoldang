package statistics

import (
	"os"
	"fmt"
	cl "github.com/OlafRadicke/banwoldang/customlogger"
)



// WriteStatistic write files with statistics
func (statistic *Statistics) WriteStatistic() {

	_, err := os.Stat(statistic.StatisticDir)
	if os.IsNotExist(err) {
		err := os.MkdirAll(statistic.StatisticDir, 0770)
		if err != nil {
			fmt.Println("error! check the log files for more information")
			cl.ErrorLogger.Fatal(err)
			return
		}
	}
	statistic.WriteDurationStatistic()
	statistic.WriteNamePartStatistic()
	statistic.WriteResolutionHeightStatistic()
	statistic.WriteResolutionWidthStatistic()
	statistic.WriteStatisticMarkdown()
	statistic.WriteTagStatistic()

}
