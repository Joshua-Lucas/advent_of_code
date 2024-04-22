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

	smallestSide := sides[0]

	for _, side := range sides {
		if side < smallestSide {
			smallestSide = side
		}
	}

	return smallestSide
}

// Calculates the ribbon need for the decorative accents of the present.
func (p Present) CalculateTheRibbonNeeded() int {
	// Find the smallest two sides. As the amount of ribbon to wrap is the
	// smallest perimeter of any one face
	smallestSide := p.Length
	secondSmallestSide := p.Width

	if p.Width < smallestSide {
		secondSmallestSide = smallestSide
		smallestSide = p.Width
	}

	if p.Height < smallestSide {
		secondSmallestSide = smallestSide
		smallestSide = p.Height
	} else if p.Height < secondSmallestSide {
		secondSmallestSide = p.Height
	}

	ribbonForWrapping := smallestSide + smallestSide + secondSmallestSide + secondSmallestSide

	// The amount of ribbon needed for the bow is equal to the volume of the package.
	ribbonForTheBow := p.Length * p.Width * p.Height

	totalRibbonNeeded := ribbonForWrapping + ribbonForTheBow

	return totalRibbonNeeded
}

// NewPresent create a present object based on the dimensions given.
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
