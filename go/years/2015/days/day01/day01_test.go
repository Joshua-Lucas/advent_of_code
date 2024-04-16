package days_test

import (
	"testing"

	days "joshualucas.dev/aoc/years/2015/days/day01"
)

type testCase struct {
	input    string
	expected int
}

var testCases = []testCase{
	{input: "(", expected: 1},
	{input: ")", expected: -1},
	{input: " ", expected: 0},
	{input: "(())", expected: 0},
	{input: "()()", expected: 0},
	{input: "(((", expected: 3},
	{input: "(()(()(", expected: 3},
	{input: "))(((((", expected: 3},
	{input: "())", expected: -1},
	{input: "))(", expected: -1},
	{input: ")))", expected: -333},
	{input: ")())())", expected: -3},
}

func TestDayOne(t *testing.T) {
	for _, tc := range testCases {
		result := days.RunDayOnePartOne(tc.input)
		if result != tc.expected {
			t.Errorf("RunDay01(%v) returns %v.Which does not result in the expect int of %v", tc.input, result, tc.expected)
		}
	}
}
