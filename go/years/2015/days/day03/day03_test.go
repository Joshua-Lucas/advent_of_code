package dayThree_test

import (
	"testing"

	dayThree "joshualucas.dev/aoc/years/2015/days/day03"
	"joshualucas.dev/aoc/years/2015/days/day03/data"
)

func TestSolvePuzzlePartOne(t *testing.T) {
	type testCaseDay struct {
		input    string
		expected int
	}

	testCases := map[string]testCaseDay{
		"did not move": {
			input:    "",
			expected: 1,
		},
		"move on x axis": {
			input:    ">",
			expected: 2,
		},
		"move all directions": {

			input:    "^>v<",
			expected: 4,
		},
		"move on y axis": {
			input:    "^v^v^v^v^v",
			expected: 2,
		},
		"do not move on wrong char": {
			input:    "\n",
			expected: 1,
		},
	}

	for name, tc := range testCases {

		t.Run(name, func(t *testing.T) {
			result, _ := dayThree.SolvePuzzle(tc.input)

			if result != tc.expected {
				t.Errorf("Result: %v, but expected %v from input: %v", result, tc.expected, tc.input)
			}
		})
	}
}

func TestSolvePuzzlePartTwo(t *testing.T) {
	type testCaseDay struct {
		input    string
		expected int
	}

	testCases := map[string]testCaseDay{
		"did not move": {
			input:    "",
			expected: 1,
		},
		"move on x axis": {
			input:    "^v",
			expected: 3,
		},
		"move all directions": {

			input:    "^>v<",
			expected: 3,
		},
		"move on y axis": {
			input:    "^v^v^v^v^v",
			expected: 11,
		},
		"do not move on wrong char": {
			input:    "\n",
			expected: 1,
		},
	}

	for name, tc := range testCases {

		t.Run(name, func(t *testing.T) {
			_, result := dayThree.SolvePuzzle(tc.input)

			if result != tc.expected {
				t.Errorf("Result: %v, but expected %v from input: %v", result, tc.expected, tc.input)
			}
		})
	}
}

func TestSliceContains(t *testing.T) {
	type inputSliceContains struct {
		slice []data.House
		value data.House
	}

	testCases := map[string]struct {
		input    inputSliceContains
		expected bool
	}{
		"Value is in slice": {
			input: inputSliceContains{
				slice: []data.House{
					{Longitude: 1, Latitude: 2},
					{Longitude: 2, Latitude: 2},
					{Longitude: 1, Latitude: 2},
					{Longitude: 1, Latitude: 3},
				},
				value: data.House{Longitude: 1, Latitude: 2},
			},
			expected: true,
		},
		"Value is not in slice": {
			input: inputSliceContains{
				slice: []data.House{
					{Longitude: 1, Latitude: 2},
					{Longitude: 2, Latitude: 2},
					{Longitude: 1, Latitude: 2},
					{Longitude: 1, Latitude: 3},
				},
				value: data.House{Longitude: 3, Latitude: 2},
			},
			expected: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := dayThree.SliceContains(tc.input.slice, tc.input.value)

			if result != tc.expected {
				t.Errorf("Result: %v, but expected %v from input: %v", result, tc.expected, tc.input)
			}

		})
	}
}
