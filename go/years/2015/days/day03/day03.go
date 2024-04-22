// Day 3: Perfectly Spherical Houses in a Vacuum
// https://adventofcode.com/2015/day/3

package dayThree

import (
	"fmt"

	"joshualucas.dev/aoc/years/2015/days/day03/data"
)

// SolvePuzzle returns part one and part two solutions for puzzle from Year 2015 Day 3 event.
func SolvePuzzle(directions string) (int, int) {

	// Create delivery log with the stating delivery of house at 0,0
	deliveryLog := []data.House{{Longitude: 0, Latitude: 0}}

	// Follow the directions given
	for _, char := range directions {

		// If there is a newline(\n) char in the directions skip it.
		if char == 10 {
			continue
		}

		//find the last house visited.
		lastHouse := deliveryLog[len(deliveryLog)-1]

		// Then, get the location of the next house to deliver at.
		houseDeliveriedAt, err := data.LogDelivery(lastHouse, string(char))

		// If there is an error with the direction we just want to skip it, and log
		// that we go an error.
		if err != nil {
			fmt.Printf("Issue with the directions. Got error: %v", err)
			continue
		}

		// Log the delivery of a present at the house.
		deliveryLog = append(deliveryLog, houseDeliveriedAt)

	}

	// Track how many house received a present.
	var atleastOnePresentDeliveried []data.House

	// check that the house was not already tracked.
	for _, house := range deliveryLog {
		if !SliceContains(atleastOnePresentDeliveried, house) {
			atleastOnePresentDeliveried = append(atleastOnePresentDeliveried, house)
		}
	}

	numberOfHousesDeliveredTo := len(atleastOnePresentDeliveried)

	return numberOfHousesDeliveredTo, 0
}

// SliceContains returns Boolean telling us if the value was in the slice.
func SliceContains(slice []data.House, value data.House) bool {
	houseMatches := false
	for _, house := range slice {
		if house.Matches(value) {
			houseMatches = true
			break
		}
	}

	return houseMatches
}
