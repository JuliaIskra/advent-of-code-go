package day_01

import (
	"bufio"
	"os"
	"strconv"
)

func Part2(filename string) (int, error) {
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
			return 0, err
		}

		for i := 0; i < diff; i++ {
			switch direction {
			case 'L':
				pointer--
				if pointer == -1 {
					pointer = 99
				}
			case 'R':
				pointer++
				if pointer == 100 {
					pointer = 0
				}
			}
			if pointer == 0 {
				zeroCounter++
			}
		}
	}
	return zeroCounter, nil
}
