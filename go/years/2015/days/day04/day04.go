// --- Day 4: The Ideal Stocking Stuffer ---
// https://adventofcode.com/2015/day/4

package dayFour

import (
	"crypto/md5"
	"strconv"
	"strings"

	hashing "joshualucas.dev/aoc/years/2015/days/day04/pkg"
)

func SolvePuzzle(instructions string) (int, int) {

	trimmedInstructions := strings.TrimRight(instructions, "\r\n")

	lowestestKeyPartOne := 1

	foundLowestKeyPartOne := false
	// Loop through positive numbers to append to the instructions

	for !foundLowestKeyPartOne {
		// Get hash value key
		key := []byte(trimmedInstructions + strconv.Itoa(lowestestKeyPartOne))

		// get md5 hash
		hash := md5.Sum(key)

		isLowestKey, _ := hashing.VerifyStartingZeros(hash, 5)

		if isLowestKey {
			foundLowestKeyPartOne = true
			break
		}

		lowestestKeyPartOne = lowestestKeyPartOne + 1

	}

	// --- PART TWO ---

	lowestestKeyPartTwo := 1
	foundLowestKeyPartTwo := false

	for !foundLowestKeyPartTwo {
		// Get hash value key
		key := []byte(trimmedInstructions + strconv.Itoa(lowestestKeyPartTwo))

		// get md5 hash
		hash := md5.Sum(key)

		isLowestKey, _ := hashing.VerifyStartingZeros(hash, 6)

		if isLowestKey {
			foundLowestKeyPartTwo = true
			break
		}

		lowestestKeyPartTwo = lowestestKeyPartTwo + 1

	}

	return lowestestKeyPartOne, lowestestKeyPartTwo

}
