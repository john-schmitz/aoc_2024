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

	for isPointInBoard(current_point, lines) {
		visited[current_point] = true
		current_point, current_direction, di = move(lines, current_point, di, directions, current_direction)
	}

	return len(visited), nil
}

func move(lines []string, current_point Point, direction_index int, directions []Point, current_direction Point) (Point, Point, int) {
	next_point := Point{current_point.i + current_direction.i, current_point.j + current_direction.j}

	if isPointInBoard(next_point, lines) && lines[next_point.i][next_point.j] == '#' {
		direction_index++
		direction_index = direction_index % 4

		current_direction = directions[direction_index]
		next_point = Point{current_point.i + current_direction.i, current_point.j + current_direction.j}

	}
	current_point = next_point
	return current_point, current_direction, direction_index
}

func isPointInBoard(point Point, board []string) bool {
	return point.i >= 0 && point.i < len(board) && point.j >= 0 && point.j < len(board[0])
}

// make a slow and fast pointer
// the move function maybe returns an error indicating that it broke boundaries.
func partTwo(filePath string) (int, error) {
	bytes, _ := os.ReadFile(filePath)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	total := 0
	fmt.Println(len(lines) * len(lines[0]))
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if lines[i][j] == '#' || lines[i][j] == '^' {
				continue
			}

			new_board := make([]string, len(lines))
			copy(new_board, lines)
			tmp := []rune(new_board[i])
			tmp[j] = '#'
			new_board[i] = string(tmp)

			if isLoop(new_board) {
				total = total + 1
			}
		}
	}

	return total, nil
}

func isLoop(lines []string) bool {
	current_point := getStartingPoint(lines)
	faster_current_point := getStartingPoint(lines)

	directions := []Point{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}

	di := 0
	d2 := 0

	current_direction := directions[0]
	faster_direction := directions[0]

	for isPointInBoard(current_point, lines) && isPointInBoard(faster_current_point, lines) {
		current_point, current_direction, di = move(lines, current_point, di, directions, current_direction)

		faster_current_point, faster_direction, d2 = move(lines, faster_current_point, d2, directions, faster_direction)
		faster_current_point, faster_direction, d2 = move(lines, faster_current_point, d2, directions, faster_direction)

		if current_point == faster_current_point {
			return true
		}
	}
	return false
}
