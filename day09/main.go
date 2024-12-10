package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partOne(filePath string) (int, error) {
	bytes, _ := os.ReadFile(filePath)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	line := lines[0]

	total_size := 0
	for _, v := range line {
		value, _ := strconv.Atoi(string(v))
		total_size = total_size + value
	}

	memory := []int{}

	current_id := 0
	for i := 0; i < len(line); i++ {
		is_file := i%2 == 0
		value, _ := strconv.Atoi(string(line[i]))
		space := -1

		if is_file {
			space = current_id
			current_id = current_id + 1
		}

		for j := 0; j < value; j++ {
			memory = append(memory, space)
		}
	}

	first_empty := -1
	last_space := -1

	for index, v := range memory {
		if v == -1 {
			first_empty = index
			break
		}
	}

	for index, v := range memory {
		if v != -1 {
			last_space = index
		}
	}

	for first_empty != last_space+1 {
		memory[last_space], memory[first_empty] = memory[first_empty], memory[last_space]
		for index, v := range memory {
			if v == -1 {
				first_empty = index
				break
			}
		}

		for index, v := range memory {
			if v != -1 {
				last_space = index
			}
		}
	}

	checksum := calculateChecksum(memory)

	return checksum, nil
}

func calculateChecksum(memory []int) int {
	checksum := 0

	for index, value := range memory {
		if value == -1 {
			continue
		}
		checksum = checksum + (index * value)
	}
	return checksum
}

func partTwo(filePath string) (int, error) {
	bytes, _ := os.ReadFile(filePath)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	line := lines[0]

	total_size := 0
	for _, v := range line {
		value, _ := strconv.Atoi(string(v))
		total_size = total_size + value
	}

	memory := []int{}

	current_id := 0
	for i := 0; i < len(line); i++ {
		is_file := i%2 == 0
		value, _ := strconv.Atoi(string(line[i]))
		space := -1

		if is_file {
			space = current_id
			current_id = current_id + 1
		}

		for j := 0; j < value; j++ {
			memory = append(memory, space)
		}
	}

	current_id--
	current_size := 0
	fmt.Println(memory, len(memory))
	for i := len(memory) - 1; i >= 0; i-- {
		for i >= 0 && current_id == memory[i] {
			current_size++
			i--
		}

		start := i + 1
		// end := start + current_size - 1
		// fmt.Println(current_id, current_size, start, end, memory[start], memory[end])

		current_empty_space_size := 0

		for j := 0; j < len(memory); j++ {
			if memory[j] == -1 {
				current_empty_space_size++
				continue
			}

			if current_size <= current_empty_space_size && start > j-current_empty_space_size {
				if current_id != memory[start] {
					panic(fmt.Errorf("current_id != memory[i]. %d != %d", current_id, memory[start]))
				}
				for z := 0; z < current_size; z++ {
					start_file := start
					start_space := j - current_empty_space_size
					memory[start_file+z], memory[start_space+z] = memory[start_space+z], memory[start_file+z]
				}
				// fmt.Println(memory)
				break
			}

			current_empty_space_size = 0
		}

		for i >= 0 && memory[i] == -1 {
			i--
		}

		current_id--
		current_size = 1
		for i >= 0 && memory[i] != current_id {
			i--
		}
	}

	checksum := calculateChecksum(memory)

	return checksum, nil
}
