package dayFive_test

import (
	"testing"

	dayFive "joshualucas.dev/aoc/years/2015/days/day05"
)

func TestSolvePuzzlePartOne(t *testing.T) {
	type testCaseDay struct {
		input    string
		expected int
	}

	testCases := map[string]testCaseDay{
		"Has all nice strings": {
			input:    "ugknbfddgicrmopn\naaa",
			expected: 2,
		},
		"Nice when two conditions overlap": {
			input:    "aaa",
			expected: 1,
		},
		"Has all Naughty strings": {
			input:    "abcdef\nhaegwjzuvuyypxyu\ndvszwmarrgswjxmb\n",
			expected: 0,
		},
		"returns the correct amount of nice strings": {
			input:    "ddugknbfddgicrmopn\naaa\nabcdef\nhaegwjzuvuyypxyu\ndvszwmarrgswjxmb\n",
			expected: 2,
		},
	}

	for name, tc := range testCases {

		t.Run(name, func(t *testing.T) {
			result, _ := dayFive.SolvePuzzle(tc.input)

			if result != tc.expected {
				t.Errorf("Result: %v, but expected %v from input: %v", result, tc.expected, tc.input)
			}
		})
	}
}
