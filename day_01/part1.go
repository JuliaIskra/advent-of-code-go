package day_01

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func Part1(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	pointer := 50
	zeroCounter := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		direction, diffStr := line[0], line[1:]
		diff, err := strconv.Atoi(diffStr)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
		switch direction {
		case 'L':
			pointer = handleOverflow(pointer - diff)
		case 'R':
			pointer = handleOverflow(pointer + diff)
		default:
			return 0, errors.New("invalid direction")
		}
		if pointer > 99 || pointer < 0 {
			return 0, errors.New("invalid pointer")
		}
		if pointer == 0 {
			zeroCounter++
		}
	}

	return zeroCounter, nil
}

func handleOverflow(pointer int) int {
	if pointer < 0 {
		return handleOverflow(100 + pointer%100)
	}
	if pointer > 99 {
		return pointer % 100
	}
	return pointer
}
