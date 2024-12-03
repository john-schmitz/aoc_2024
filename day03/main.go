package main

import (
	"fmt"
	"os"
	"strings"
)

func partOne(filePath string) (int, error) {
	bytes, _ := os.ReadFile(filePath)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	prefix := "mul("
	total := 0
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			if i+len(prefix) > len(line) {
				break
			}
			x := line[i : i+len(prefix)]
			if x != prefix {
				continue
			}

			for j := i + 6; j < i+12; j++ {
				if line[j] == ')' {
					result, err := parseMul(line[i : j+1])
					if err != nil {
						break
					}
					total = total + result
					break
				}
			}
		}
	}

	return total, nil
}

func parseMul(line string) (int, error) {
	if len(line) > 12 {
		panic("len(line) > 12")
	}
	var l int
	var r int
	_, err := fmt.Sscanf(line, "mul(%d,%d)", &l, &r)

	return l * r, err
}

func partTwo(filePath string) (int, error) {
	bytes, _ := os.ReadFile(filePath)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	isMulEnabled := true

	prefix := "mul("
	doPrefix := "do()"
	dontPrefix := "don't()"

	total := 0
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			if i+len(doPrefix) > len(line) {
				break
			}
			y := line[i : i+len(doPrefix)]
			if y == doPrefix {
				isMulEnabled = true
				continue
			}

			if i+len(dontPrefix) > len(line) {
				break
			}
			z := line[i : i+len(dontPrefix)]
			if z == dontPrefix {
				isMulEnabled = false
				continue
			}

			if !isMulEnabled {
				continue
			}

			if i+len(prefix) > len(line) {
				break
			}
			x := line[i : i+len(prefix)]
			if x != prefix {
				continue
			}

			for j := i + 6; j < i+12; j++ {
				if line[j] == ')' {
					result, err := parseMul(line[i : j+1])
					if err != nil {
						break
					}
					total = total + result
					break
				}
			}
		}
	}

	return total, nil
}
