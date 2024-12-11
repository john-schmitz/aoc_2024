package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func blinkStones(stones []int) []int {
	result := []int{}

	for _, stone := range stones {
		blink_result := blinkStone(stone)
		result = append(result, blink_result...)
	}

	return result
}

func blinkStone(stone int) []int {
	if stone == 0 {
		return []int{1}
	}

	x := fmt.Sprint(stone)
	if len(x)%2 == 0 {
		left := x[:len(x)/2]
		right := x[(len(x) / 2):]

		l, _ := strconv.Atoi(left)
		r, _ := strconv.Atoi(right)

		return []int{l, r}
	}

	return []int{stone * 2024}
}

func partOne(filePath string) (int, error) {
	bytes, _ := os.ReadFile(filePath)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	line := lines[0]

	stones := []int{}

	for _, v := range strings.Split(line, " ") {
		num, _ := strconv.Atoi(v)
		stones = append(stones, num)
	}

	i := 0

	for i < 25 {
		stones = blinkStones(stones)
		i++
	}

	return len(stones), nil
}

func partTwo(filePath string) (int, error) {
	bytes, _ := os.ReadFile(filePath)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	line := lines[0]

	stones := []int{}

	for _, v := range strings.Split(line, " ") {
		num, _ := strconv.Atoi(v)
		stones = append(stones, num)
	}

	i := 0

	for i < 75 {
		stones = blinkStones(stones)
		i++
	}

	return len(stones), nil
}
