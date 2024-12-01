package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func partOne(filePath string) int {
	lines, err := getLines(filePath)
	if err != nil {
		panic(err)
	}
	left, right, err := parseLines(lines)
	if err != nil {
		panic(err)
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	total := 0

	for i := 0; i < len(left); i++ {
		x := Abs(left[i] - right[i])
		total += x
	}

	return total
}

func partTwo(filePath string) int {
	lines, err := getLines(filePath)
	if err != nil {
		panic(err)
	}
	left, right, err := parseLines(lines)
	if err != nil {
		panic(err)
	}

	total := 0

	for i := 0; i < len(left); i++ {
		occ := 0
		for j := 0; j < len(right); j++ {
			if left[i] == right[j] {
				occ += 1
			}
		}

		total += left[i] * occ
	}

	return total
}

func parseLines(lines []string) ([]int, []int, error) {
	left := make([]int, 0)
	right := make([]int, 0)
	for _, s := range lines {
		result := strings.Split(s, "   ")
		value, err := strconv.Atoi(result[0])
		if err != nil {
			return nil, nil, err
		}
		left = append(left, value)

		value, err = strconv.Atoi(result[1])
		if err != nil {
			return nil, nil, err
		}

		right = append(right, value)
	}

	return left, right, nil
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
