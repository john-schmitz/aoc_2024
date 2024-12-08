package main

import (
	"os"
	"strconv"
	"strings"
)

func partOne(filePath string) (int, error) {
	bytes, _ := os.ReadFile(filePath)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	total := 0

	for _, line := range lines {
		splited := strings.Split(line, ":")
		target, _ := strconv.Atoi(splited[0])

		x := strings.Split(strings.TrimSpace(splited[1]), " ")

		values := make([]int, len(x))
		for i := 0; i < len(x); i++ {
			v, _ := strconv.Atoi(x[i])
			values[i] = v
		}

		if isValidTotal(values[1:], target, values[0]) {
			total = total + target
		}
	}

	return total, nil
}

func isValidTotal(numbers []int, target int, acc int) bool {
	if len(numbers) == 0 {
		return acc == target
	}

	if acc > target {
		return false
	}

	return isValidTotal(numbers[1:], target, acc+numbers[0]) || isValidTotal(numbers[1:], target, acc*numbers[0])
}

func isValidTotal2(numbers []int, target int, acc int) bool {
	if len(numbers) == 0 {
		return acc == target
	}

	if acc > target {
		return false
	}

	x := strconv.Itoa(acc) + strconv.Itoa(numbers[0])
	value, err := strconv.Atoi(x)

	if err != nil {
		panic(err)
	}

	return isValidTotal2(numbers[1:], target, acc+numbers[0]) ||
		isValidTotal2(numbers[1:], target, acc*numbers[0]) ||
		isValidTotal2(numbers[1:], target, value)
}

func partTwo(filePath string) (int, error) {
	bytes, _ := os.ReadFile(filePath)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	total := 0

	for _, line := range lines {
		splited := strings.Split(line, ":")
		target, _ := strconv.Atoi(splited[0])

		x := strings.Split(strings.TrimSpace(splited[1]), " ")

		values := make([]int, len(x))
		for i := 0; i < len(x); i++ {
			v, _ := strconv.Atoi(x[i])
			values[i] = v
		}

		if isValidTotal2(values[1:], target, values[0]) {
			total = total + target
		}
	}

	return total, nil
}
