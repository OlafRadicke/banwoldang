package statistics

import (
	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// ContResolutionHeight counts the founded resolution heights.
func (statistic *Statistics) ContResolutionHeight(height int) {
	if count, ok := statistic.ResolutionHeight[height]; ok {
		cl.InfoLogger.Println("The height ", height, " is allready added (", count, "). Count up...")
		statistic.ResolutionHeight[height]++
	} else {
		cl.InfoLogger.Println("The height ", height, " (", count, ") is new.")
		statistic.ResolutionHeight[height] = 1
	}
}
