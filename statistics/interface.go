package statistics

// statistics The inteface of the feletree struckt
type statistics interface {
	GoThroughCollection()
}

type Statistics struct {
	UsedTags     map[string]int
	FoundedFiles int
}

// Statistics create new instance of Statistics and get it back.
func NewStatistics() Statistics {
	statistics := Statistics{}

	statistics.FoundedFiles = 0
	statistics.UsedTags = make(map[string]int)
	return statistics
}
