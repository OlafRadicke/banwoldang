package statistics


// AddCheckSumLocations Add a checksum and a location.
func (statistic *Statistics) AddCheckSumLocations(checksum string, path string) {
	statistic.CheckSumLocations[checksum] = append(statistic.CheckSumLocations[checksum], path )
}
