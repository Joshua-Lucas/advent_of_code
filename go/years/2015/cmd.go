package years

import (
	"fmt"
	"os"

	inputUtils "joshualucas.dev/aoc/utils"
	days "joshualucas.dev/aoc/years/2015/days/day01"
)

func RunYear2015(day int) {

	rootPath, _ := os.Getwd()

	inputsPath := rootPath + "/years/2015/days"

	switch day {
	case 1:

		dayPath := "/day01/input.txt"

		instructions, err := inputUtils.ReadInputFile(inputsPath + dayPath)

		if err != nil {
			fmt.Printf("Error with reading input file. %v", err)
			return
		}

		partOneAnswer := days.RunDayOnePartOne(instructions)

		fmt.Printf("The anser to day one puzzle is: %v", partOneAnswer)

	default:
		fmt.Println("There is not an answer to that day of the puzzle")
	}
}
