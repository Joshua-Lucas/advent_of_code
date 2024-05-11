package dayFour_test

import (
	"testing"

	dayFour "joshualucas.dev/aoc/years/2015/days/day04"
)

func TestSolvePuzzlePartOne(t *testing.T) {
	type testCaseDay struct {
		input    string
		expected int
	}

	testCases := map[string]testCaseDay{
		"1 Has 5 zeros": {
			input:    "abcdef",
			expected: 609043,
		},
		"2 Has 5 zeros": {
			input:    "pqrstuv",
			expected: 1048970,
		},
		"Does not pass": {
			input:    "abcdef",
			expected: 1203,
		},
	}

	for name, tc := range testCases {

		t.Run(name, func(t *testing.T) {
			result, _ := dayFour.SolvePuzzle(tc.input)

			if result != tc.expected {
				t.Errorf("Result: %v, but expected %v from input: %v", result, tc.expected, tc.input)
			}
		})
	}
}
