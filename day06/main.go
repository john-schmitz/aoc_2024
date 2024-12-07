package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	i int
	j int
}

func getStartingPoint(lines []string) Point {
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if lines[i][j] == '^' {
				return Point{i, j}
			}
		}
	}

	panic(fmt.Errorf("starting point not found for input"))
}

func partOne(filePath string) (int, error) {
	bytes, _ := os.ReadFile(filePath)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	x := len(lines)
	y := len(lines[0])
	current_point := getStartingPoint(lines)
	directions := []Point{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}

	di := 0

	current_direction := directions[0]

	visited := make(map[Point]bool)

	for current_point.i >= 0 && current_point.i < x && current_point.j >= 0 && current_point.j < y {
		current_point, current_direction, di = move(lines, current_point, di, directions, current_direction)
		visited[current_point] = true
		next_point := Point{current_point.i + current_direction.i, current_point.j + current_direction.j}
		current_point = next_point
	}

	return len(visited), nil
}

func move(lines []string, current_point Point, di int, directions []Point, current_direction Point) (Point, Point, int) {
	if lines[current_point.i][current_point.j] == '#' {
		di++
		if di == len(directions) {
			di = 0
		}
		previous_point := Point{current_point.i - current_direction.i, current_point.j - current_direction.j}
		current_direction = directions[di]
		current_point = previous_point
	}
	return current_point, current_direction, di
}

// make a slow and fast pointer
// the move function maybe returns an error indicating that it broke boundaries.
func partTwo(filePath string) (int, error) {
	bytes, _ := os.ReadFile(filePath)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	x := len(lines)
	y := len(lines[0])
	current_point := getStartingPoint(lines)
	directions := []Point{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}

	di := 0

	current_direction := directions[0]

	visited := make(map[Point]bool)

	for current_point.i >= 0 && current_point.i < x && current_point.j >= 0 && current_point.j < y {
		current_point, current_direction, di = move(lines, current_point, di, directions, current_direction)
		visited[current_point] = true
		next_point := Point{current_point.i + current_direction.i, current_point.j + current_direction.j}
		current_point = next_point
	}

	return len(visited), nil
}
