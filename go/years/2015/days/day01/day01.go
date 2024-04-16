package days

// Runs the logic to solve day one
func RunDayOnePartOne(instructions string) int {

	// Starting floor
	currentFloor := 0

	// loop through input characters up count.
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
