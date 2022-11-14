package statistics

import (
	"errors"
	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

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

