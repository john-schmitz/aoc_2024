package main

import (
	"os"
	"strings"
)

func partOne(filePath string) (int, error) {
	bytes, _ := os.ReadFile(filePath)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	xmas := "XMAS"
	directions := []Point{
		{1, 0},   // UP
		{-1, 0},  // DOWN
		{0, 1},   // RIGHT
		{0, -1},  // LEFT
		{1, 1},   // NE
		{-1, 1},  // SE
		{-1, -1}, // SW
		{1, -1},  // NW
	}

	total := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if lines[i][j] == 'X' {
				for _, direction := range directions {
					for z := 1; z < len(xmas); z++ {
						x := i + direction.i*z
						y := j + direction.j*z
						if x < 0 || y < 0 || x >= len(lines) || y >= len(lines[0]) {
							break
						}

						next := lines[x][y]

						if next != xmas[z] {
							break
						}

						if next == 'S' {
							total += 1
						}
					}

				}
			}
		}
	}

	return total, nil
}

type Point struct {
	i int
	j int
}

func partTwo(filePath string) (int, error) {
	bytes, _ := os.ReadFile(filePath)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	directions := []Point{
		{1, 1},   // NE
		{-1, 1},  // SE
		{-1, -1}, // SW
		{1, -1},  // NW
	}

	total := 0

	for i := 0; i < len(lines); i++ {
	outer:
		for j := 0; j < len(lines[0]); j++ {
			if lines[i][j] == 'A' && (1+i >= 0 && j+1 >= 0 && 1+i < len(lines) && j+1 < len(lines[0])) && (-1+i >= 0 && j+1 >= 0 && -1+i < len(lines) && j+1 < len(lines[0])) && (-1+i >= 0 && j+-1 >= 0 && -1+i < len(lines) && j+-1 < len(lines[0])) && (1+i >= 0 && j+-1 >= 0 && 1+i < len(lines) && j+-1 < len(lines[0])) {
				for _, direction := range directions {
					if lines[direction.i+i][direction.j+j] != 'M' && lines[direction.i+i][direction.j+j] != 'S' {
						continue outer
					}
				}

				if lines[directions[0].i+i][directions[0].j+j] == lines[directions[2].i+i][directions[2].j+j] {
					continue outer
				}

				if lines[directions[1].i+i][directions[1].j+j] == lines[directions[3].i+i][directions[3].j+j] {
					continue outer
				}

				total += 1
			}
		}
	}

	return total, nil
}
