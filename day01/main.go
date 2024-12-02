package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func partOne(filePath string) int {
	bytes, _ := os.ReadFile(filePath)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	n := len(lines)

	left := make([]int, n)
	right := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Sscanf(lines[i], "%d %d", &left[i], &right[i])
	}

	slices.Sort(left)
	slices.Sort(right)

	sum := 0
	for index := 0; index < n; index++ {
		diff := left[index] - right[index]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}

	return sum
}

func partTwo(filePath string) int {
	bytes, _ := os.ReadFile(filePath)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	n := len(lines)

	left := make([]int, n)
	right := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Sscanf(lines[i], "%d %d", &left[i], &right[i])
	}

	sum := 0

	for index := 0; index < len(left); index++ {
		occ := 0
		for jindex := 0; jindex < len(right); jindex++ {
			if left[index] == right[jindex] {
				occ += 1
			}
		}

		sum += left[index] * occ
	}

	return sum
}
