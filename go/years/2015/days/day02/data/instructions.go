package data

import (
	"fmt"
	"strconv"
	"strings"
)

// Takes in the instruction string and splits all the instructions into present
// structs so the array of presents can be itterated through
func PrepInstructions(rawInstructions string) []Present {
	var presentInstructions []Present

	// Take instructions and split on new line character. The reson we use FieldFunc
	// vs split is becuase Split would leave and empty string at the end of the slice
	rawStringPresents := strings.FieldsFunc(rawInstructions, func(r rune) bool {
		return r == '\n'
	})

	// Then take new slice and split on "x"
	for _, presentDimensions := range rawStringPresents {
		// This will store [length, width, height]
		dimensionsArray := strings.Split(presentDimensions, "x")

		// convert each sting in the dimenstionArray to a int value.
		length, err := strconv.Atoi(dimensionsArray[0])

		if err != nil {
			fmt.Printf("Error with length(%v) string to int conversion:%v", dimensionsArray[0], err)
		}

		width, err := strconv.Atoi(dimensionsArray[1])

		if err != nil {
			fmt.Println("Error with width string to int conversion", err)
		}

		height, err := strconv.Atoi(dimensionsArray[2])

		if err != nil {
			fmt.Println("Error with height string to int conversion", err)
		}

		// Create a new Present struct for the present demensions.
		present := NewPresent(length, width, height)

		// Add the Present to an array of presents.
		presentInstructions = append(presentInstructions, present)
	}

	return presentInstructions
}
