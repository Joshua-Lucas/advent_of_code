package days_test

import (
	"testing"

	days "joshualucas.dev/aoc/years/2015/days/day01"
)

type testCase struct {
	input    string
	expected int
}

var testCasesPartOne = []testCase{
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
	{input: ")))", expected: -3},
	{input: ")())())", expected: -3},
}

func TestDayOnePartOne(t *testing.T) {
	for _, tc := range testCasesPartOne {
		result := days.RunDayOnePartOne(tc.input)
		if result != tc.expected {
			t.Errorf("RunDayOnePartOne(%v) returns %v.Which does not result in the expect int of %v", tc.input, result, tc.expected)
		}
	}
}

var testCasesPartTwo = []testCase{
	{input: "(((", expected: 0},
	{input: ")", expected: 1},
	{input: "()())", expected: 5},
}

func TestDayOnePartTwo(t *testing.T) {
	for _, tc := range testCasesPartTwo {
		result := days.RunDayOnePartTwo(tc.input)
		if result != tc.expected {
			t.Errorf("RunDayOnePartTwo(%v) returns %v.Which does not result in the expect int of %v", tc.input, result, tc.expected)
		}
	}
}
