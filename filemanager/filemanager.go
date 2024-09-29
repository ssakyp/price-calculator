package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("Failed to open file.")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// this can be used to identify the error
	scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Failed to read file!")
	}

	file.Close()
	return lines, nil
}
