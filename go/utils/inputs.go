package inputUtils

import "os"

// ReadInputFile returns a string value of the content of the file input.txt file
func ReadInputFile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(content), nil

}
