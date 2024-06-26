// Solution for Advent to Code day 2 of year 2015
// Puzzle can be found at https://adventofcode.com/2015/day/2

package dayTwo

import (
	"joshualucas.dev/aoc/years/2015/days/day02/data"
)

func SolvePuzzle(instructions string) (int, int) {
	totalWrappingPaper := 0

	// Prep instructions to iterate over to calculate total wrapping paper.
	parsedInstructions := data.PrepInstructions(instructions)

	// iterate over the instructions and calculate total wrapping needed.
	for _, present := range parsedInstructions {
		wrappingNeededForPresent := present.SurfaceArea + present.SmallestSide

		totalWrappingPaper = totalWrappingPaper + wrappingNeededForPresent
	}

	// Part Two
	totalRibbonNeeded := 0

	// iterate over the present list and calculate total ribbon needed
	for _, present := range parsedInstructions {
		ribbonForPresent := present.CalculateTheRibbonNeeded()

		totalRibbonNeeded = totalRibbonNeeded + ribbonForPresent
	}

	return totalWrappingPaper, totalRibbonNeeded
}
