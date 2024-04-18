package dayOne

// Runs the logic to solve day one
func RunDayOnePartOne(instructions string) int {

	// Starting floor
	currentFloor := 0

	// loop through input characters .
	for _, char := range instructions {
		if char == '(' {
			currentFloor = currentFloor + 1
		}

		if char == ')' {
			currentFloor = currentFloor - 1
		}

	}

	return currentFloor
}

func RunDayOnePartTwo(instructions string) int {

	var positionEnteredBasement int
	currentPosition := 0
	currentFloor := 0

	for _, char := range instructions {
		currentPosition = currentPosition + 1

		if char == '(' {
			currentFloor = currentFloor + 1
		}

		if char == ')' {
			currentFloor = currentFloor - 1
		}

		// When user goes into the basement.
		if currentFloor < 0 {
			positionEnteredBasement = currentPosition
			break

		}
	}

	return positionEnteredBasement

}
