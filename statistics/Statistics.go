package statistics

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

	// A list of Duplicates with an list of his locations
	CheckSumLocations map[string][]string
}

