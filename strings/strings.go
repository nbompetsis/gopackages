package strings

import (
	"bufio"
	"os"
)

func ReadLinesFromFile(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, err
}
