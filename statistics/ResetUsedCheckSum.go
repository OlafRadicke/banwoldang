package statistics

// ResetUsedCheckSum make a reset of stored checksums.
func (statistic *Statistics) ResetUsedCheckSum() () {
	statistic.UsedCheckSum = make(map[string]int)
}