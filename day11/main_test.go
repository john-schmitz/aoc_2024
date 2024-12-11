package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBlinkStones(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{input: []int{125, 17}, expected: []int{253000, 1, 7}},
		{input: []int{253000, 1, 7}, expected: []int{253, 0, 2024, 14168}},
		{input: []int{253, 0, 2024, 14168}, expected: []int{512072, 1, 20, 24, 28676032}},
		{input: []int{512072, 1, 20, 24, 28676032}, expected: []int{512, 72, 2024, 2, 0, 2, 4, 2867, 6032}},
	}

	for _, tC := range testCases {
		t.Run(fmt.Sprint(tC.input), func(t *testing.T) {
			result := blinkStones(tC.input)

			if !reflect.DeepEqual(result, tC.expected) {
				t.Errorf("Expected blinkStones(%v) = %v. Got: %v", tC.input, tC.expected, result)
			}
		})
	}
}

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
