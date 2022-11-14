package statistics

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
	statistic.UsedCheckSum = make(map[string]int)
	return &statistic
}