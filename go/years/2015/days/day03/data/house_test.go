package data_test

import (
	"testing"

	"joshualucas.dev/aoc/years/2015/days/day03/data"
)

func TestHouseMatch(t *testing.T) {
	testCases := map[string]struct {
		input    [2]data.House
		expected bool
	}{

		"Houses match": {
			input: [2]data.House{
				{Longitude: 1, Latitude: 2},
				{Longitude: 1, Latitude: 2},
			},
			expected: true,
		},
    "Houses match: negative int": {
			input: [2]data.House{
				{Longitude: -1, Latitude: -2},
				{Longitude: -1, Latitude: -2},
			},
			expected: true,
		},
		"Houses do not match: flip long and lat": {
			input: [2]data.House{
				{Longitude: 1, Latitude: 2},
				{Longitude: 2, Latitude: 1},
			},
			expected: false,
		},
		"House does not match: diff lat": {
			input: [2]data.House{
				{Longitude: 1, Latitude: 0},
				{Longitude: 1, Latitude: 2},
			},
			expected: false,
		},
		"House does not match: diff long": {
			input: [2]data.House{
				{Longitude: 9, Latitude: 2},
				{Longitude: 1, Latitude: 2},
			},
			expected: false,
		},
		"House does not match: negative int": {
			input: [2]data.House{
				{Longitude: -1, Latitude: -2},
				{Longitude: 1, Latitude: 2},
			},
			expected: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			house1 := tc.input[0]
			result := house1.Matches(tc.input[1])
			if result != tc.expected {
				t.Errorf("The current house: %v does not return the expect bool of %v with the address given of %v", house1, tc.expected, tc.input[1])
			}
		})
	}
}
