package main

import (
	"os"
	"strings"
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
			expected: 480,
		},
		{
			desc:     "input",
			filepath: "input.txt",
			expected: 25629,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			bytes, _ := os.ReadFile(tC.filepath)
			lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

			result, err := partOne(lines)
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
			expected: 25629,
		},
		{
			desc:     "input",
			filepath: "input.txt",
			expected: -1,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			bytes, _ := os.ReadFile(tC.filepath)
			lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

			result, err := partTwo(lines)
			if err != nil {
				panic(err)
			}
			if result != tC.expected {
				t.Errorf("Expected partTwo(%s) = %d. But got %d", tC.filepath, tC.expected, result)
			}
		})
	}
}
