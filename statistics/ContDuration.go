package statistics

import (
	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// ContDuration counts the founded uration (in ).
func (statistic *Statistics) ContDuration(duration int) {
	if count, ok := statistic.Duration[duration]; ok {
		cl.InfoLogger.Println("The duration ", duration, " is allready added (", count, "). Count up...")
		statistic.Duration[duration]++
	} else {
		cl.InfoLogger.Println("The duration ", duration, " (", count, ") is new.")
		statistic.Duration[duration] = 1
	}
}

