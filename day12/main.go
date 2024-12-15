package main

import (
	"os"
	"strings"
)

type Point struct {
	x int
	y int
}

type Region = map[Point]bool

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
	directions := []Point{
		{1, 0},  // UP
		{0, 1},  // RIGHT
		{-1, 0}, // DOWN
		{0, -1}, // LEFT
	}

	x := len(lines)
	y := len(lines[0])

	reached := make(map[Point]bool)
	regions := make([]Region, 0)

	total_price := 0

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if reached[Point{i, j}] {
				continue
			}

			region := make(Region)

			frontier := make([]Point, 1)
			frontier[0] = Point{i, j}

			reached[Point{i, j}] = true
			region[Point{i, j}] = true
			crop := lines[i][j]

			for len(frontier) > 0 {
				current := frontier[0]
				frontier = frontier[1:]

				for _, direction := range directions {
					next_point := direction.sum(current)

					if inBounds(next_point, lines) && lines[next_point.x][next_point.y] == crop && !region[next_point] {
						frontier = append(frontier, next_point)
						region[next_point] = true
						reached[next_point] = true
					}
				}
			}

			regions = append(regions, region)
		}
	}

	for _, region := range regions {
		perimeter := 0

		for point := range region {
			for _, direction := range directions {
				next_point := direction.sum(point)

				if !inBounds(next_point, lines) || lines[next_point.x][next_point.y] != lines[point.x][point.y] {
					perimeter++
				}
			}
		}

		total_price += len(region) * perimeter
	}

	return total_price, nil
}

func partTwo(filePath string) (int, error) {
	bytes, _ := os.ReadFile(filePath)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	return len(lines), nil
}
