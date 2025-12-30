package day_03

import (
	"bufio"
	"os"
	"strconv"
)

func Part1(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		tens, ind := findMaxN(line, 0, len(line)-1)
		digit, _ := findMaxN(line, ind+1, len(line))
		sum += tens*10 + digit
	}

	return sum, nil
}

func findMaxN(line string, start int, end int) (int, int) {
	maxN := 0
	ind := -1
	for i := start; i < end; i++ {
		n, _ := strconv.Atoi(string(line[i]))
		if n > maxN {
			maxN = n
			ind = i
		}
	}
	return maxN, ind
}
