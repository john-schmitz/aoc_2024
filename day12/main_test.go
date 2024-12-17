package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestBoardPrice(t *testing.T) {
	testCases := []struct {
		input    []string
		expected int
	}{
		{
			input:    []string{"AAAA"},
			expected: 40,
		},
		{
			input:    []string{"AAAA", "BBBC"},
			expected: 68,
		},
		{
			input:    []string{"AAAA", "BBCD", "BBCC", "EEEC"},
			expected: 140,
		},
	}

	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%v", tC.input), func(t *testing.T) {
			result, err := partOne(tC.input)
			if err != nil {
				panic(err)
			}

			if result != tC.expected {
				t.Errorf("Expected partOne(%s) = %d. But got %d", tC.input, tC.expected, result)
			}
		})
	}
}

func TestPartOne(t *testing.T) {
	testCases := []struct {
		desc     string
		filepath string
		expected int
	}{
		{
			desc:     "sample",
			filepath: "sample.txt",
			expected: 1930,
		},
		{
			desc:     "input",
			filepath: "input.txt",
			expected: 1396562,
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
			expected: 1206,
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
