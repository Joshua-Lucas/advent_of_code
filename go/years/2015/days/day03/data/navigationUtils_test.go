package data_test

import (
	"testing"

	"joshualucas.dev/aoc/years/2015/days/day03/data"
)

func TestLogDelivery(t *testing.T) {

	type inputCaseLogDelivery struct {
		ref data.House
		dir string
	}

	testCases := map[string]struct {
		input    inputCaseLogDelivery
		expected data.House
	}{

		"log delivery to the east(>)": {
			input:    inputCaseLogDelivery{data.House{0, 0}, ">"},
			expected: data.House{1, 0},
		},
		"log deliver to the west(<)": {
			input:    inputCaseLogDelivery{data.House{0, 0}, "<"},
			expected: data.House{-1, 0},
		},
		"log return to start": {
			input:    inputCaseLogDelivery{data.House{1, 0}, "<"},
			expected: data.House{0, 0},
		},
		"log delivery to the north(^)": {
			input:    inputCaseLogDelivery{data.House{0, 0}, "^"},
			expected: data.House{0, 1},
		},
		"log delivery to the south(v)": {
			input:    inputCaseLogDelivery{data.House{1, 0}, "v"},
			expected: data.House{1, -1},
		},
		"Do not move on invalid direction": {
			input:    inputCaseLogDelivery{data.House{1, 0}, "\n"},
			expected: data.House{1, 0},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result, err := data.LogDelivery(tc.input.ref, tc.input.dir)

			// Test that error occurs and same location is returned
			if err != nil {
				if !result.Matches(tc.expected) {
					t.Errorf("LogDelivery(%v) returned %v which was not the expect result of %v", tc.input, result, tc.expected)
				}
			}

			if err == nil {
				if result != tc.expected {
					t.Errorf("LogDelivery(%v) returned %v which was not the expect result of %v", tc.input, result, tc.expected)
				}

			}
		})
	}
}
