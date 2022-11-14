package statistics

import (
	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

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

