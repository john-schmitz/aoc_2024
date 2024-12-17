package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type GameCoordinates struct {
	ButtonA Point
	ButtonB Point
	Prize   Point
}

func parseCoordinates(input []string) ([]GameCoordinates, error) {
	var result []GameCoordinates
	var current GameCoordinates
	lineCount := 0

	// Regular expression to match coordinates
	re := regexp.MustCompile(`X[\+=](\d+),\s*Y[\+=](\d+)`)

	for _, line := range input {
		if line == "" {
			continue
		}

		// Extract coordinates using regex
		matches := re.FindStringSubmatch(line)
		if len(matches) != 3 {
			continue
		}

		// Parse X and Y values
		x, err := strconv.Atoi(matches[1])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(matches[2])
		if err != nil {
			return nil, err
		}

		// Determine which point to update based on the line content
		switch {
		case strings.HasPrefix(line, "Button A:"):
			current.ButtonA = Point{x: int(x), y: int(y)}
			lineCount++
		case strings.HasPrefix(line, "Button B:"):
			current.ButtonB = Point{x: int(x), y: int(y)}
			lineCount++
		case strings.HasPrefix(line, "Prize:"):
			current.Prize = Point{x: int(x), y: int(y)}
			lineCount++
		}

		// If we've collected all three points, add to result
		if lineCount == 3 {
			result = append(result, current)
			current = GameCoordinates{} // Reset for next set
			lineCount = 0
		}
	}

	return result, nil
}

func partOne(lines []string) (int, error) {
	total := 0
	games, err := parseCoordinates(lines)

	if err != nil {
		panic(err)
	}

	for _, game := range games {
		a := game.ButtonA
		b := game.ButtonB
		prize := game.Prize
		ca := float64(prize.x*b.y-prize.y*b.x) / float64(a.x*b.y-a.y*b.x)
		cb := (float64(prize.x) - float64(a.x)*ca) / float64(b.x)

		if math.Mod(ca, 1) == 0 && math.Mod(cb, 1) == 0 {
			total += int(ca*3 + cb)
		}
	}

	return total, nil
}

func partTwo(lines []string) (int, error) {
	total := 0
	games, err := parseCoordinates(lines)

	if err != nil {
		panic(err)
	}

	for _, game := range games {
		a := game.ButtonA
		b := game.ButtonB
		prize := game.Prize
		prize.x = prize.x + 10000000000000
		prize.y = prize.y + 10000000000000
		ca := float64(prize.x*b.y-prize.y*b.x) / float64(a.x*b.y-a.y*b.x)
		cb := (float64(prize.x) - float64(a.x)*ca) / float64(b.x)

		if math.Mod(ca, 1) == 0 && math.Mod(cb, 1) == 0 {
			total += int(ca*3 + cb)
		}
	}

	return total, nil
}
