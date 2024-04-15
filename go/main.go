package main

import (
	"flag"
	"fmt"

	years "joshualucas.dev/aoc/years/2015"
)

// Declare varaibles that will hold the flags passed when the file is run.
var yearFlag *int
var dayFlag *int

func init() {
	// On init grab the flags passed when running the program.
	yearFlag = flag.Int("year", 2105, "The year that the puzzle is fron")

	dayFlag = flag.Int("day", 1, "The day of the puzzle")

	flag.Parse()
}

func main() {

	switch *yearFlag {
	case 2015:
		years.RunYear2015(*dayFlag)
	default:
		fmt.Println("Invalid year selected, please run again with a valid year")
	}
}
