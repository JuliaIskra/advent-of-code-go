package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
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
			return
		}
		switch direction {
		case 'L':
			pointer = handleOverflow(pointer - diff)
		case 'R':
			pointer = handleOverflow(pointer + diff)
		default:
			fmt.Println("Invalid direction")
			return
		}
		if pointer > 99 || pointer < 0 {
			fmt.Println("Invalid pointer", pointer)
			return
		}
		if pointer == 0 {
			zeroCounter++
		}
	}

	fmt.Println(zeroCounter)
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
