package day_03

import (
	"bufio"
	"math"
	"os"
)

func Part2(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		counter := 11
		start := 0
		end := len(line) - counter
		for counter >= 0 {
			tens, ind := findMaxN(line, start, end)
			start = ind + 1
			end++
			sum = tens*int(math.Pow(10, float64(counter))) + sum
			counter--
		}
	}

	return sum, nil
}
