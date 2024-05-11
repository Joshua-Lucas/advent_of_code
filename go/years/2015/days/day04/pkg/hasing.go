package hashing

import (
	"encoding/hex"
)

// Returns a boolean if the hash start with the amount of zeros passes as the second argument.
func VerifyStartingZeros(hash [16]byte, amountOfZeros int) (bool, error) {

	encodedHash := hex.EncodeToString(hash[:])

	var matchingChar []bool

	for idx, char := range encodedHash {
		if idx > amountOfZeros-1 {
			break
		}

		if char == '0' {
			matchingChar = append(matchingChar, true)
		}

	}

	if len(matchingChar) == amountOfZeros {
		return true, nil
	}

	return false, nil
}
