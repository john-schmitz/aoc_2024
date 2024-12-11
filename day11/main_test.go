package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBlinkStone(t *testing.T) {
	testCases := []struct {
		input    int
		expected []int
	}{
		{input: 0, expected: []int{1}},
		{input: 1, expected: []int{2024}},
		{input: 1000, expected: []int{10, 0}},
	}

	for _, tC := range testCases {
		t.Run(fmt.Sprint(tC.input), func(t *testing.T) {
			result := blinkStone(tC.input)

			if !reflect.DeepEqual(result, tC.expected) {
				t.Errorf("Expected blinkStone(%v) = %v. Got: %v", tC.input, tC.expected, result)
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
			expected: 55312,
		},
		{
			desc:     "input",
			filepath: "input.txt",
			expected: 213625,
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
			expected: 65601038650482,
		},
		{
			desc:     "input",
			filepath: "input.txt",
			expected: 252442982856820,
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
