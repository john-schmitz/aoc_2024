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
			expected: 36,
		},
		{
			desc:     "input",
			filepath: "input.txt",
			expected: 531,
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
			expected: 81,
		},
		{
			desc:     "input",
			filepath: "input.txt",
			expected: 1210,
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
