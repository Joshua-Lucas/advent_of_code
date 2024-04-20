package data_test

import (
	"testing"

	"joshualucas.dev/aoc/years/2015/days/day02/data"
)

type testCaseParseInstructions struct {
	input    string
	expected data.Present
}

var testCasesParseInstructions = []testCaseParseInstructions{
	{
		input: "2x3x4\n1x1x10",
		expected: data.Present{
			Length:       2,
			Width:        3,
			Height:       4,
			SurfaceArea:  52,
			SmallestSide: 6,
		},
	},
	{
		input: "1x1x10\n2x3x4",
		expected: data.Present{
			Length:       1,
			Width:        1,
			Height:       10,
			SurfaceArea:  42,
			SmallestSide: 1,
		},
	},
}

func TestPrepInstructions(t *testing.T) {
	for _, tc := range testCasesParseInstructions {

		result := data.PrepInstructions(tc.input)

		if result[0].Length != tc.expected.Length {
			t.Errorf("NewPresent(%v) returns a value %v for key Length.Which does not result[0] in the expect Length of %v", tc.input, result[0].Length, tc.expected.Length)
		}

		if result[0].Height != tc.expected.Height {
			t.Errorf("NewPresent(%v) returns a value %v for key Height.Which does not result[0] in the expect Height of %v", tc.input, result[0].Height, tc.expected.Height)
		}

		if result[0].Width != tc.expected.Width {
			t.Errorf("NewPresent(%v) returns a value %v for key Width.Which does not result[0] in the expect Width of %v", tc.input, result[0].Width, tc.expected.Width)
		}

		if result[0].SurfaceArea != tc.expected.SurfaceArea {
			t.Errorf("NewPresent(%v) returns a value %v for key SurfaceArea.Which does not result[0] in the expect SurfaceArea of %v", tc.input, result[0].SurfaceArea, tc.expected.SurfaceArea)
		}

		if result[0].SmallestSide != tc.expected.SmallestSide {
			t.Errorf("NewPresent(%v) returns a value %v for key SmallestSide.Which does not result[0] in the expect SmallestSide of %v", tc.input, result[0].SmallestSide, tc.expected.SmallestSide)
		}
	}
}
