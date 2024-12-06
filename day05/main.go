package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partOne(filePath string) (int, error) {
	bytes, _ := os.ReadFile(filePath)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	rules := make(map[int][]int)

	var i int
	for i = 0; i < len(lines); i++ {
		line := lines[i]

		if line == "" {
			break
		}

		var l int
		var r int

		fmt.Sscanf(line, "%d|%d", &l, &r)

		rule, ok := rules[l]

		if !ok {
			rule = make([]int, 1)
		}

		rule = append(rule, r)

		rules[l] = rule
	}

	total := 0

	for i = i + 1; i < len(lines); i++ {
		line := strings.Split(lines[i], ",")
		update := make([]int, len(line))

		for j := 0; j < len(line); j++ {
			n, err := strconv.Atoi(line[j])

			if err != nil {
				panic(err)
			}
			update[j] = n
		}

		if !isValidUpdate(update, rules) {
			continue
		}

		mid := len(update) / 2

		total += update[mid]

	}
	return total, nil
}

func partTwo(filePath string) (int, error) {
	bytes, _ := os.ReadFile(filePath)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	rules := make(map[int][]int)

	var i int
	for i = 0; i < len(lines); i++ {
		line := lines[i]

		if line == "" {
			break
		}

		var l int
		var r int

		fmt.Sscanf(line, "%d|%d", &l, &r)

		rule, ok := rules[l]

		if !ok {
			rule = make([]int, 1)
		}

		rule = append(rule, r)

		rules[l] = rule
	}

	total := 0

	for i = i + 1; i < len(lines); i++ {
		line := strings.Split(lines[i], ",")
		update := make([]int, len(line))
		for j := 0; j < len(line); j++ {
			n, err := strconv.Atoi(line[j])

			if err != nil {
				panic(err)
			}
			update[j] = n
		}

		if isValidUpdate(update, rules) {
			continue
		}

		valid := transformInvalidUpdate(update, rules)
		mid := len(valid) / 2

		total += valid[mid]
	}

	return total, nil
}

func transformInvalidUpdate(update []int, rules map[int][]int) []int {
	result := make([]int, len(update))
	copy(result, update)

	for i := len(result) - 1; i > 0; i-- {
		n := result[i]
		rule := rules[n]

	outer:
		for j := i - 1; j >= 0; j-- {
			for _, v := range rule {
				if result[j] != v {
					continue
				}

				from := j
				to := i
				result = shiftSlice(result, from, to)
				i++
				break outer
			}
		}
	}

	if !isValidUpdate(result, rules) {
		fmt.Println(fmt.Errorf("invalid update\n before: %v\n after: %v", update, result))
	}

	return result
}

func isValidUpdate(update []int, rules map[int][]int) bool {
	for i := len(update) - 1; i > 0; i-- {
		n := update[i]
		rule := rules[n]

		for j := i - 1; j >= 0; j-- {
			for _, v := range rule {
				if update[j] == v {
					return false
				}
			}
		}
	}

	return true
}

func shiftSlice(input []int, from, to int) []int {
	if from >= len(input) || to >= len(input) || from < 0 || to < 0 {
		panic(fmt.Errorf("invalid input for slice of length %d. Got from: %d to: %d", len(input), from, to))
	}

	result := make([]int, len(input))
	copy(result, input)

	last := result[from]

	for i := from; i < to; i++ {
		result[i] = result[i+1]
	}

	result[to] = last

	return result
}
