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

		fmt.Println("Is invalid", update)

		for z := len(update) - 1; z > 0; z-- {
			n := update[z]
			rule := rules[n]

		outer:
			for j := z - 1; j >= 0; j-- {
				for _, v := range rule {
					if update[j] != v {
						continue
					}
					new_update := make([]int, len(update))
					copy(new_update, update)

					new_update[z], new_update[j] = new_update[j], new_update[z]

					if isValidUpdate(new_update, rules) {
						mid := len(new_update) / 2
						fmt.Println(update, new_update, new_update[mid])

						total += new_update[mid]
						break outer
					}
				}
			}
		}

	}

	return total, nil
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
