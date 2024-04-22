package data

import (
	"errors"
	"fmt"
)

// LogDelivery returns the house the was just delivered to.
func LogDelivery(lastLocation House, direction string) (House, error) {
	// Set the new house location to the last location by default, as we will
	// only move a single direction at a time.
	newHouseLong := lastLocation.Longitude
	newHouseLat := lastLocation.Latitude

	switch direction {
	case ">":
		newHouseLong = newHouseLong + 1
	case "<":
		newHouseLong = newHouseLong - 1
	case "^":
		newHouseLat = newHouseLat + 1
	case "v":
		newHouseLat = newHouseLat - 1
	default:
		return lastLocation, errors.New(fmt.Sprintf("%v is an invalid direction", direction))
	}

	return House{
		Longitude: newHouseLong,
		Latitude:  newHouseLat,
	}, nil
}
