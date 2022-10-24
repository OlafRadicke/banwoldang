package statistics

import (
	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// statistics The inteface of the feletree struckt
type statistics interface {
	ContPartsOfNames(string)
	GoThroughCollection()
	WriteDurationStatistic()
	WriteNamePartStatistic()
	WriteResolutionHeightStatistic()
	WriteResolutionWidthStatistic()
	WriteStatistic()
	WriteStatisticMarkdown()
	WriteTagStatistic()
}

type Statistics struct {
	StatisticDir     string
	UsedTags         map[string]int
	PartsOfNames     map[string]int
	Duration         map[int]int
	ResolutionHeight map[int]int
	ResolutionWidth  map[int]int
	FoundedFiles     int
}

// NewStatistics create new instance of Statistics and get it back.
func NewStatistics(statisticDir string) *Statistics {
	statistic := Statistics{}
	statistic.StatisticDir = statisticDir
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
