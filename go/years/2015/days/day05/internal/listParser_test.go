package listParser_test

import (
	"testing"

	listParser "joshualucas.dev/aoc/years/2015/days/day05/internal"
)

func TestContainsThreeVowels(t *testing.T) {
	type testCaseDay struct {
		input    string
		expected bool
	}

	testCases := map[string]testCaseDay{
		"Has three vowels": {
			input:    "abcdefghi",
			expected: true,
		},
		"Has only vowels": {
			input:    "aaa",
			expected: true,
		},
		"Has some vowels but not enough": {
			input:    "afdgrett",
			expected: false,
		},
		"Has no vowels": {
			input:    "pyhwtjklh",
			expected: false,
		},
	}

	for name, tc := range testCases {

		t.Run(name, func(t *testing.T) {
			result := listParser.ContainsThreeVowels(tc.input)

			if result != tc.expected {
				t.Errorf("Result: %v, but expected %v from input: %v", result, tc.expected, tc.input)
			}
		})
	}
}

func TestContainsForbidden(t *testing.T) {
	type testCaseDay struct {
		input    string
		expected bool
	}

	testCases := map[string]testCaseDay{
		"Contains Forbidden ab + cd": {
			input:    "abcdef",
			expected: true,
		},
		"Contains forbidden cd": {
			input:    "cdroeg",
			expected: true,
		},
		"Contains forbidden pq": {
			input:    "afdpqfter",
			expected: true,
		},
		"Contains forbiddend xy": {
			input:    "pyhwtjklh",
			expected: true,
		},
		"Forbidden similar but out of order xy": {
			input:    "yxtuv",
			expected: false,
		},
		"Does not contain a pair": {
			input:    "hjlkjasdfg",
			expected: false,
		},
	}

	for name, tc := range testCases {

		t.Run(name, func(t *testing.T) {
			result := listParser.ContainsThreeVowels(tc.input)

			if result != tc.expected {
				t.Errorf("Result: %v, but expected %v from input: %v", result, tc.expected, tc.input)
			}
		})
	}
}
