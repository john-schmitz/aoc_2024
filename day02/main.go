package main

import (
	"fmt"
	"os"
	"strings"
)

func partOne(filePath string) (int, error) {
	bytes, _ := os.ReadFile(filePath)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	n := len(lines)

	lists := make([][]int, n)

	for i := 0; i < n; i++ {
		numbers := strings.Split(lines[i], " ")
		lists[i] = make([]int, len(numbers))
		for j := 0; j < len(numbers); j++ {
			fmt.Sscanf(numbers[j], "%d", &lists[i][j])
		}
	}

	total := 0
	for i := 0; i < n; i++ {
		if isSafe(lists[i]) {
			total += 1
		}
	}

	return total, nil
}

func partTwo(filePath string) (int, error) {
	bytes, _ := os.ReadFile(filePath)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	n := len(lines)

	lists := make([][]int, n)

	for i := 0; i < n; i++ {
		numbers := strings.Split(lines[i], " ")
		lists[i] = make([]int, len(numbers))
		for j := 0; j < len(numbers); j++ {
			fmt.Sscanf(numbers[j], "%d", &lists[i][j])
		}
	}

	total := 0
	for i := 0; i < n; i++ {
		if isSafe(lists[i]) {
			total += 1
			continue
		}
		for j := 0; j < len(lists[i]); j++ {
			if isSafe(removeIndex(lists[i], j)) {
				total += 1
				break
			}
		}
	}

	return total, nil
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func isSafe(input []int) bool {
	var l int

	increasing := input[0]-input[1] < 0
	if input[0]-input[1] == 0 {
		return false
	}

	for j := 0; j < len(input); j++ {
		curr := input[j]

		if l == 0 {
			l = curr
			continue
		}

		d := curr - l

		if d == 0 {
			return false
		}

		if increasing && d < 0 {
			return false
		}

		if !increasing && d > 0 {
			return false
		}

		if d < 0 {
			d = -d
		}

		if d > 3 {
			return false
		}

		l = curr
	}

	return true
}
