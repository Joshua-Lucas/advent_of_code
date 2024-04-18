package data

type Present struct {
	Length       int
	Width        int
	Height       int
	SurfaceArea  int
	SmallestSide int
}

// utility function used to get the smallest side of a present.
func FindSmallestSide(a int, b int, c int) int {
	sides := [3]int{a, b, c}

	// Code have used a simplier solution but wanted to write an alg as I was taking
	// and alg course at the time.
	for i := 0; i < len(sides); i++ {
		for j := 0; j < len(sides)-1-i; j++ {
			if sides[j] > sides[j+1] {
				temp := sides[j]
				sides[j] = sides[j+1]
				sides[j+1] = temp
			}
		}
	}

	return sides[0]
}

// Create a factory function to create package struct
func NewPresent(l int, w int, h int) Present {
	// All Pesents are right rectangle prism so we can use this calculation
	// 2(l*w) + 2(w*h) + 2(h*l) to find the total suface area of the present.
	areaOfSideA := l * w
	areaOfSideB := w * h
	areaOfSideC := h * l

	// calculate the surface area of the whole presenet.
	sufaceAreaOfPresent := (2 * areaOfSideA) + (2 * areaOfSideB) + (2 * areaOfSideC)

	// find the smallets side based on the areas
	smallestSide := FindSmallestSide(areaOfSideA, areaOfSideB, areaOfSideC)

	return Present{
		Length:       l,
		Width:        w,
		Height:       h,
		SurfaceArea:  sufaceAreaOfPresent,
		SmallestSide: smallestSide,
	}
}
