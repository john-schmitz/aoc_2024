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
			expected: 161,
		},
		{
			desc:     "input",
			filepath: "input.txt",
			expected: 174960292,
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
			expected: 48,
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

func TestParseMul(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected int
	}{
		{
			desc:     "parseMul",
			input:    "mul(982,733)",
			expected: 719806,
		},
		{
			desc:     "parseMul",
			input:    "mul(308,100)",
			expected: 30800,
		},
		{
			desc:     "parseMul",
			input:    "mul(539,64)",
			expected: 34496,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := parseMul(tC.input)

			if err != nil {
				t.Error(err)
			}

			if result != tC.expected {
				t.Errorf("Expected parseMul(%s) = %d. But got %d", tC.input, tC.expected, result)
			}
		})
	}
}
