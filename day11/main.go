package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func blinkStones(stones []int, iterations int) int {
	stonesMap := make(map[int]int)

	for _, stone := range stones {
		stonesMap[stone] = 1
	}

	for i := 0; i < iterations; i++ {
		newMap := make(map[int]int)

		for stone, quantity := range stonesMap {
			for _, transformed_stone := range blinkStone(stone) {
				current, _ := newMap[transformed_stone]

				newMap[transformed_stone] = current + quantity
			}
		}
		stonesMap = newMap
	}

	total := 0
	for _, v := range stonesMap {
		total += v
	}

	return total
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

	return blinkStones(stones, 25), nil
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

	return blinkStones(stones, 75), nil
}
