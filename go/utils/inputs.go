package inputUtils

import "os"

func ReadInputFile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(content), nil

}
