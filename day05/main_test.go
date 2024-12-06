package main

import (
	"reflect"
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
			expected: 143,
		},
		{
			desc:     "input",
			filepath: "input.txt",
			expected: 6034,
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
			expected: 123,
		},
		{
			desc:     "sample",
			filepath: "input.txt",
			expected: 6305,
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

func TestMoveSlice(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []int
		from     int
		to       int
		expected []int
	}{
		{
			desc:     "swap elements",
			input:    []int{1, 2},
			expected: []int{2, 1},
			from:     0,
			to:       1,
		},
		{
			desc:     "swap elements",
			input:    []int{1, 2, 3},
			expected: []int{2, 3, 1},
			from:     0,
			to:       2,
		},
		{
			desc:     "shift right",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 3, 4, 2, 5},
			from:     1,
			to:       3,
		},
		{
			desc:     "shift right",
			input:    []int{97, 13, 75, 47, 29},
			expected: []int{97, 75, 47, 29, 13},
			from:     1,
			to:       4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := shiftSlice(tC.input, tC.from, tC.to)

			if !reflect.DeepEqual(result, tC.expected) {
				t.Errorf("Invalid output. Expected %v got %v", tC.expected, result)
			}
		})
	}
}
