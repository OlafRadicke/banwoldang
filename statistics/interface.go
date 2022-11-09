package statistics

// statistics The inteface of the feletree struckt
type statistics interface {
	ContDuration(int)
	ContPartsOfNames(string)
	ContResolutionHeight(int)
	ContResolutionWidth(int)
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
	// Count how often files has a duration (in minutes).
	Duration map[int]int
	// Count of fonded media files
	FoundedFiles int
	// Count how often file name parts is in using
	PartsOfNames map[string]int
	// Count how often height of resolution is in using
	ResolutionHeight map[int]int
	// Count how often width of resolution is in using
	ResolutionWidth map[int]int
	// The direktory of the statistic files
	StatisticDir string
	// Count how often tags is in using
	UsedTags map[string]int
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
