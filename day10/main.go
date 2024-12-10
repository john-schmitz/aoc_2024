package main

import (
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

/*
...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9
*/

func partOne(filePath string) (int, error) {
	bytes, _ := os.ReadFile(filePath)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	directions := []Point{
		{1, 0},  // UP
		{0, 1},  // RIGHT
		{-1, 0}, // DOWN
		{0, -1}, // LEFT
	}

	trails := getTrails(lines)
	total := 0

	for _, trail := range trails {
		reached := make(map[Point]bool)
		frontier := make([]Point, 1)

		frontier[0] = trail
		reached[trail] = true

		for len(frontier) > 0 {
			current := frontier[0]

			if lines[current.x][current.y] == '9' {
				total++
			}

			frontier = frontier[1:]
			neighbors := []Point{}

			for _, direction := range directions {
				next_point := Point{direction.x + current.x, direction.y + current.y}
				if inBounds(next_point, lines) && canMove(current, next_point, lines) {
					neighbors = append(neighbors, next_point)
				}
			}

			for _, neighbor := range neighbors {
				_, ok := reached[neighbor]
				if !ok {
					frontier = append(frontier, neighbor)
					reached[neighbor] = true
				}
			}
		}

	}

	return total, nil
}

func canMove(from, to Point, board []string) bool {
	l, _ := strconv.Atoi(string(board[from.x][from.y]))
	r, _ := strconv.Atoi(string(board[to.x][to.y]))

	return l == r-1
}

func inBounds(point Point, board []string) bool {
	return point.x >= 0 && point.x < len(board) &&
		point.y >= 0 && point.y < len(board[0])
}

func getTrails(lines []string) []Point {
	trails := []Point{}

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if lines[i][j] == '0' {
				trails = append(trails, Point{i, j})
			}
		}
	}

	return trails
}

func partTwo(filePath string) (int, error) {
	bytes, _ := os.ReadFile(filePath)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	return len(lines), nil
}
