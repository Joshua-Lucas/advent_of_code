package years

import (
	"fmt"
	"os"

	inputUtils "joshualucas.dev/aoc/utils"
	dayOne "joshualucas.dev/aoc/years/2015/days/day01"
	dayTwo "joshualucas.dev/aoc/years/2015/days/day02"
)

/*
	RunYear2015 is a function that runs the solutions for the puzzles of the

specified day in the year 2015. It takes an integer representing the day of
the puzzle to be run.
*/
func RunYear2015(day int) {

	rootPath, _ := os.Getwd()

	// Declare the general path for all inputs for the 2015 package.
	inputsPath := rootPath + "/years/2015/days"

	// Add a switch case to process the day input and run the proper function.
	switch day {
	case 1:

		dayPath := "/day01/input.txt"

		instructions, err := inputUtils.ReadInputFile(inputsPath + dayPath)

		if err != nil {
			fmt.Printf("Error with reading input file. %v", err)
			return
		}

		// Get the answers for the puzzle.
		partOneAnswer := dayOne.RunDayOnePartOne(instructions)

		fmt.Printf("The answer to part one of the day one puzzle is: %v \n", partOneAnswer)

		partTwoAnswer := dayOne.RunDayOnePartTwo(instructions)

		fmt.Printf("The anwser to part two of the day one puzzle is: %v \n", partTwoAnswer)
	case 2:

		dayPath := "/day02/input.txt"

		instructions, err := inputUtils.ReadInputFile(inputsPath + dayPath)

		if err != nil {
			fmt.Printf("Error with reading input file. %v", err)
			return
		}

		partOne, partTwo := dayTwo.SolvePuzzle(instructions)

		fmt.Printf("The answer to part one of the day two puzzle is: %v \n", partOne)

		fmt.Printf("The anwser to part two of the day two puzzle is: %v \n", partTwo)

	default:
		fmt.Println("There is not an answer to that day of the puzzle")
	}
}
