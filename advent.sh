#!/bin/bash

# Check if a day parameter was provided
if [ $# -ne 1 ]; then
    echo "Usage: $0 <day>"
    exit 1
fi

# Format day to have leading zero if needed
day=$(printf "%02d" $1)

# Create directory
folder_name="day$day"
mkdir -p "$folder_name"
cd "$folder_name"

# Initialize Go module
go mod init "$folder_name"

# Create main.go with basic template
cat > main.go << 'EOL'
package main

import (
	"bufio"
	"os"
)

func partOne(filePath string) (int, error) {
	return 0, nil
}

func partTwo(filePath string) (int, error) {
	return 0, nil
}

func getLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
EOL

# Create main_test.go with basic test template
cat > main_test.go << 'EOL'
package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	testCases := []struct {
		desc     string
		filepath string
		expected int
	}{
		{
			desc:     "sample",
			filepath: "sample.txt",
			expected: -1,
		},
		{
			desc:     "input",
			filepath: "input.txt",
			expected: -1,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := partOne(tC.filepath)
			if err != nil {
				panic(err)
			}

			if result != tC.expected {
				t.Errorf("Expected partOne(%s) = %d. But got %d", tC.filepath, tC.expected, result)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	testCases := []struct {
		desc     string
		filepath string
		expected int
	}{
		{
			desc:     "sample",
			filepath: "sample.txt",
			expected: -1,
		},
		{
			desc:     "input",
			filepath: "input.txt",
			expected: -1,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := partTwo(tC.filepath)
			if err != nil {
				panic(err)
			}
			if result != tC.expected {
				t.Errorf("Expected partTwo(%s) = %d. But got %d", tC.filepath, tC.expected, result)
			}
		})
	}
}
EOL

echo "Created $folder_name with Go module and starter files"