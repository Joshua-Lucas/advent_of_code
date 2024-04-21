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
	{input: "2x3x4\n2x3x4\n", expected: 116},
	{input: "1x1x10\n1x1x10\n", expected: 86},
}

func TestDayTwoPartOne(t *testing.T) {
	for _, tc := range testCasesPartOne {
		result, _ := dayTwo.RunPartOne(tc.input)
		if result != tc.expected {
			t.Errorf("RunDayTwo(%v) returns %v.Which does not result in the expect int of %v", tc.input, result, tc.expected)
		}
	}
}

var testCasesPartTwo = []testCaseDay{
	{input: "2x3x4\n3x4x2\n", expected: 68},
	{input: "1x1x10\n10x1x1\n", expected: 28},
}

func TestDayTwoPartTwo(t *testing.T) {
	for _, tc := range testCasesPartTwo {
		_, result := dayTwo.RunPartOne(tc.input)
		if result != tc.expected {
			t.Errorf("RunDayTwo(%v) returns %v.Which does not result in the expect int of %v", tc.input, result, tc.expected)
		}
	}
}
