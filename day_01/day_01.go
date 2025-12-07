package day_01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Part1(filename string) {
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

func Part2(filename string) {
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
		var p, zeros int
		switch direction {
		case 'L':
			p, zeros = CountZeros(pointer, -diff)
		case 'R':
			p, zeros = CountZeros(pointer, diff)
		default:
			fmt.Println("Invalid direction")
			return
		}
		println(pointer, line, p, zeros, zeroCounter)
		pointer = p
		zeroCounter += zeros
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

func CountZeros(oldPointer int, diff int) (int, int) {
	newPointer := oldPointer + diff
	zeroCounter := 0
	if oldPointer > 0 && newPointer < 0 || oldPointer < 0 && newPointer > 0 {
		zeroCounter++
	}

	if newPointer == 100 {
		return 0, zeroCounter
	}
	if newPointer < 0 {
		mod, _ := CountZeros(100+newPointer%100, 0)
		return mod, -newPointer/100 + zeroCounter
	}
	if newPointer > 99 {
		return newPointer % 100, newPointer/100 + zeroCounter
	}
	return newPointer, 0
}
