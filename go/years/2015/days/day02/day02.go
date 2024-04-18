package dayTwo

import (
	"fmt"

	"joshualucas.dev/aoc/years/2015/days/day02/data"
)

func RunPartOne(instructions string) int {

	// figure out how to take the string and turn break it down into different gifts.

	// Create a struct that that holds the gifts l,w,h,smallest side, and surface area
	// loop through gift list instructions creating an array of structs
	// then loop through that array of gift structs to calculate the need amount of wrapping paper
	// return that value

	// or
	aPresent := data.NewPresent(2, 3, 4)
	// create a value to store total gift surface area
	// create a function that returns surface needed per gift
	// return needed amount of wrapping paper.
	fmt.Println(aPresent)
	return 1
}
