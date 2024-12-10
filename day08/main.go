package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Point struct {
	x int
	y int
}

func partOne(filePath string) (int, error) {
	bytes, _ := os.ReadFile(filePath)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	targets := make(map[byte][]Point)

	visited := make(map[Point]bool)

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			current_point := Point{i, j}
			value := lines[current_point.x][current_point.y]
			if value != '.' && lines[current_point.x][current_point.y] != '#' {
				a, ok := targets[value]
				if !ok {
					a = make([]Point, 0)
				}

				a = append(a, current_point)
				targets[value] = a
			}
		}
	}
	for _, points := range targets {
		for _, pair := range generatePointPairs(points) {
			A := pair[0]
			B := pair[1]

			V := Point{B.x - A.x, B.y - A.y}
			last := Point{A.x - V.x, A.y - V.y}
			first := Point{B.x + V.x, B.y + V.y}

			if isInBoard(last, lines) {
				visited[last] = true
			}

			if isInBoard(first, lines) {
				visited[first] = true
			}

		}
	}

	return len(visited), nil
}

func partTwo(filePath string) (int, error) {
	bytes, _ := os.ReadFile(filePath)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	targets := make(map[byte][]Point)

	visited := make(map[Point]bool)

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			current_point := Point{i, j}
			value := lines[current_point.x][current_point.y]
			if value != '.' && lines[current_point.x][current_point.y] != '#' {
				a, ok := targets[value]
				if !ok {
					a = make([]Point, 0)
				}

				a = append(a, current_point)
				targets[value] = a
			}
		}
	}
	for _, points := range targets {
		for _, pair := range generatePointPairs(points) {
			A := pair[0]
			B := pair[1]
			V := Point{B.x - A.x, B.y - A.y}

			current := A
			for isInBoard(current, lines) {
				visited[current] = true
				current = Point{current.x + V.x, current.y + V.y}
			}

			current = Point{A.x - V.x, A.y - V.y}
			for isInBoard(current, lines) {
				visited[current] = true
				current = Point{current.x - V.x, current.y - V.y}
			}
		}
	}

	return len(visited), nil
}

func isInBoard(point Point, board []string) bool {
	return 0 <= point.x && point.x < len(board) && 0 <= point.y && point.y < len(board[0])
}

func isCollinear(p1, p2, p3 Point) bool {
	area := p1.x*(p2.y-p3.y) + p2.x*(p3.y-p1.y) + p3.x*(p1.y-p2.y)
	return area == 0
}

func manhattanDistance(p1, p2 Point) int {
	dx := int(math.Abs(float64(p1.x - p2.x)))
	dy := int(math.Abs(float64(p1.y - p2.y)))
	return dx + dy
}

func generatePointPairs(points []Point) [][2]Point {
	pairs := make([][2]Point, 0, (len(points)*(len(points)-1))/2)

	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			pairs = append(pairs, [2]Point{points[i], points[j]})
		}
	}

	return pairs
}

func main() {
	bytes, _ := os.ReadFile("sample.txt")
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	targets := make(map[byte][]Point)

	visited := make(map[Point]bool)

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			current_point := Point{i, j}
			value := lines[current_point.x][current_point.y]
			if value != '.' && lines[current_point.x][current_point.y] != '#' {
				a, ok := targets[value]
				if !ok {
					a = make([]Point, 0)
				}

				a = append(a, current_point)
				targets[value] = a
			}
		}
	}
	fmt.Println(targets)
	for _, points := range targets {
		for _, pair := range generatePointPairs(points) {
			A := pair[0]
			B := pair[1]

			for i := 0; i < len(lines); i++ {
				for j := 0; j < len(lines[0]); j++ {
					current_point := Point{i, j}
					if isCollinear(A, B, current_point) {

						if manhattanDistance(A, current_point) == manhattanDistance(B, current_point)*2 {
							visited[current_point] = true
						} else if manhattanDistance(B, current_point) == manhattanDistance(A, current_point)*2 {
							visited[current_point] = true
						}
					}
				}
			}
		}
	}

	fmt.Println(len(visited))

}
