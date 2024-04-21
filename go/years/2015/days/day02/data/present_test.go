package data_test

import (
	"testing"

	"joshualucas.dev/aoc/years/2015/days/day02/data"
)

type testCaseArray struct {
	input    [3]int
	expected int
}

var testCasesSmallestSide = []testCaseArray{
	{input: [3]int{6, 12, 8}, expected: 6},
	{input: [3]int{1, 1, 10}, expected: 1},
}

func TestSmallestSide(t *testing.T) {
	for _, tc := range testCasesSmallestSide {
		result := data.FindSmallestSide(tc.input[0], tc.input[1], tc.input[2])
		if result != tc.expected {
			t.Errorf("smallestSide(%v) returns %v.Which does not result in the expect int of %v", tc.input, result, tc.expected)
		}
	}
}

type testCasePresent struct {
	input    [3]int
	expected data.Present
}

var testCasesPresents = []testCasePresent{
	{
		input: [3]int{2, 3, 4},
		expected: data.Present{
			Length:       2,
			Width:        3,
			Height:       4,
			SurfaceArea:  52,
			SmallestSide: 6,
		},
	},
	{
		input: [3]int{1, 1, 10},
		expected: data.Present{
			Length:       1,
			Width:        1,
			Height:       10,
			SurfaceArea:  42,
			SmallestSide: 1,
		},
	},
}

func TestPresentFactory(t *testing.T) {
	for _, tc := range testCasesPresents {
		var result data.Present
		result = data.NewPresent(tc.input[0], tc.input[1], tc.input[2])

		if result.Length != tc.expected.Length {
			t.Errorf("NewPresent(%v) returns a value %v for key Length.Which does not result in the expect Length of %v", tc.input, result.Length, tc.expected.Length)
		}

		if result.Height != tc.expected.Height {
			t.Errorf("NewPresent(%v) returns a value %v for key Height.Which does not result in the expect Height of %v", tc.input, result.Height, tc.expected.Height)
		}

		if result.Width != tc.expected.Width {
			t.Errorf("NewPresent(%v) returns a value %v for key Width.Which does not result in the expect Width of %v", tc.input, result.Width, tc.expected.Width)
		}

		if result.SurfaceArea != tc.expected.SurfaceArea {
			t.Errorf("NewPresent(%v) returns a value %v for key SurfaceArea.Which does not result in the expect SurfaceArea of %v", tc.input, result.SurfaceArea, tc.expected.SurfaceArea)
		}

		if result.SmallestSide != tc.expected.SmallestSide {
			t.Errorf("NewPresent(%v) returns a value %v for key SmallestSide.Which does not result in the expect SmallestSide of %v", tc.input, result.SmallestSide, tc.expected.SmallestSide)
		}
	}
}

type testCaseRibbonNeeded struct {
	input    [3]int
	expected int
}

var testCaseRibbonsNeeded = []testCaseRibbonNeeded{
	{input: [3]int{2, 3, 4}, expected: 34},
	{input: [3]int{1, 1, 10}, expected: 14},
	{input: [3]int{10, 1, 1}, expected: 14},
	{input: [3]int{4, 3, 2}, expected: 34},
	{input: [3]int{8, 9, 7}, expected: 534},
}

func TestCalculateRibbonNeeded(t *testing.T) {
	for _, tc := range testCaseRibbonsNeeded {
		present := data.NewPresent(tc.input[0], tc.input[1], tc.input[2])

		result := present.CalculateTheRibbonNeeded()

		if result != tc.expected {
			t.Errorf("The CalculateTheRibbonNeeded method on %v returns a value of %v which does not match the expected value of %v", present, result, tc.expected)
		}
	}

}
