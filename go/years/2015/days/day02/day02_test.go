package dayTwo_test

import (
	"testing"

	dayTwo "joshualucas.dev/aoc/years/2015/days/day02"
)

type testCaseDay struct {
	input    string
	expected int
}

var testCasesPartOne = []testCaseDay{
	{input: "2x3x4", expected: 58},
	{input: "1x1x10", expected: 43},
}

func TestDayTwoPartOne(t *testing.T) {
	for _, tc := range testCasesPartOne {
		result := dayTwo.RunPartOne(tc.input)
		if result != tc.expected {
			t.Errorf("RunDayTwo(%v) returns %v.Which does not result in the expect int of %v", tc.input, result, tc.expected)
		}
	}
}
