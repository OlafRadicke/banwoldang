package statistics

import (
	"errors"
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


// ContUsedCheckSum count part of file names.
func (statistic *Statistics) ContUsedCheckSum(checksum string) (error) {
	if count, ok := statistic.UsedCheckSum[checksum]; ok {
		cl.ErrorLogger.Println("The check sum ", checksum, " is allready added (", count, "). Count up...")
		statistic.UsedCheckSum[checksum]++
		return errors.New("duplicate checksum")
	} else {
		cl.InfoLogger.Println("The check sum ", checksum, " is new.")
		statistic.UsedCheckSum[checksum] = 1
		return nil
	}
}