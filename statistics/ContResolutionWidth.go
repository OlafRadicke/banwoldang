package statistics

import (
	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// ContResolutionWidth counts the founded resolution heights.
func (statistic *Statistics) ContResolutionWidth(width int) {
	if count, ok := statistic.ResolutionWidth[width]; ok {
		cl.InfoLogger.Println("The width ", width, " is allready added (", count, "). Count up...")
		statistic.ResolutionWidth[width]++
	} else {
		cl.InfoLogger.Println("The width ", width, " (", count, ") is new.")
		statistic.ResolutionWidth[width] = 1
	}
}
