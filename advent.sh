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
	"fmt"
)

func solve(input string) int {
	return 0
}

func main() {
	result := solve("")
	fmt.Printf("Result: %v\n", result)
}
EOL

# Create main_test.go with basic test template
cat > main_test.go << 'EOL'
package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "example 1",
			input:    "",
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solve(tc.input)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}
EOL

echo "Created $folder_name with Go module and starter files"