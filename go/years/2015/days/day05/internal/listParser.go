package listParser

// Returns true when the string contains at least three vowels
func ContainsThreeVowels(str string) bool {

	var vowelCount int

	vowelSet := [5]rune{'a', 'e', 'i', 'o', 'u'}

	for _, char := range str {

		// loop through vowel set to see if it contains a vowel.
		for _, vowel := range vowelSet {
			if vowel == char {
				vowelCount = vowelCount + 1
				break
			}
		}
	}

	return vowelCount >= 3
}

// Returns true when at least one letter appears twice in a row(i.e. aa,bb)
func ALetterAppearsTwiceInARow(str string) bool {

	return false
}

// Returns true when string contains ab, cd, pq, xy
func ContainForbiddenPairs(str string) bool {

	return false
}
