package main

import (
	"os"
	"strings"
)

type Point struct {
	x int
	y int
}

func (p Point) sum(point Point) Point {
	return Point{p.x + point.x, p.y + point.y}
}

func inBounds(point Point, board []string) bool {
	return point.x >= 0 && point.x < len(board) &&
		point.y >= 0 && point.y < len(board[0])
}

func nextUnreachedPoint(x int, y int, reached map[Point]bool) (Point, bool) {
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if !reached[Point{i, j}] {
				return Point{i, j}, true
			}
		}
	}

	return Point{x, y}, false
}

func partOne(lines []string) (int, error) {
	reached := make(map[Point]bool)
	frontier := make([]Point, 1)

	frontier[0] = Point{0, 0}
	current_cursor := lines[0][0]

	directions := []Point{
		{1, 0},  // UP
		{0, 1},  // RIGHT
		{-1, 0}, // DOWN
		{0, -1}, // LEFT
	}

	x := len(lines)
	y := len(lines[0])

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			reached[Point{i, j}] = false
		}
	}

	reached[Point{0, 0}] = true
	current_area := 0
	current_perimether := 0

	total_price := 0
	for len(frontier) > 0 {
		current := frontier[0]
		frontier = frontier[1:]
		for _, direction := range directions {
			next_point := direction.sum(current)

			if !inBounds(next_point, lines) || lines[next_point.x][next_point.y] != current_cursor {
				current_perimether++
			}

			if inBounds(next_point, lines) && lines[next_point.x][next_point.y] == current_cursor && !reached[next_point] {
				frontier = append(frontier, next_point)
				reached[next_point] = true
			}
		}

		current_area++
		reached[current] = true

		if len(frontier) == 0 {
			next, found := nextUnreachedPoint(x, y, reached)
			if found {
				frontier = append(frontier, next)
				current_cursor = lines[next.x][next.y]
			}
			total_price += current_area * current_perimether
			current_area = 0
			current_perimether = 0
		}
	}

	return total_price, nil
}

func partTwo(filePath string) (int, error) {
	bytes, _ := os.ReadFile(filePath)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	return len(lines), nil
}
