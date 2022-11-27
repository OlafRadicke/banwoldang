package statistics

import (
	"errors"
)

// AddCheckSumLocations Add a checksum and a location.
func (statistic *Statistics) AddCheckSumLocations(checksum string, path string) {
	statistic.CheckSumLocations[checksum] = append(statistic.CheckSumLocations[checksum], path )
}

// AddCheckSumLocations Add a checksum and a location.
func (statistic *Statistics) GetCheckSumLocationsCount(checksum string) (int, error) {
	if _, ok := statistic.CheckSumLocations[checksum]; ok {
		return len(statistic.CheckSumLocations[checksum]), nil
	} else {
		return 0, errors.New("hash no found")
	}
}