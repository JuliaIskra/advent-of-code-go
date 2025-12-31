package main

import (
	"advent-of-code-go/day_01"
	"advent-of-code-go/day_02"
	"advent-of-code-go/day_03"
	"advent-of-code-go/day_04"
	"advent-of-code-go/day_08"
	"advent-of-code-go/day_10"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run ./cmd <day> <part>")
		os.Exit(2)
	}

	dayStr := os.Args[1]
	part := os.Args[2]

	dayNum, err := strconv.Atoi(dayStr)
	if err != nil || dayNum <= 0 {
		fmt.Println("Invalid day, expected a positive integer")
		os.Exit(2)
	}

	// Normalize folder name like day_01, day_02
	inputPath := fmt.Sprintf("day_%02d/input.txt", dayNum)

	var res int

	switch dayNum {
	case 1:
		switch part {
		case "1":
			res, err = day_01.Part1(inputPath)
		case "2":
			res, err = day_01.Part2(inputPath)
		default:
			fmt.Println("Invalid part. Use 1 or 2.")
			os.Exit(2)
		}
	case 2:
		switch part {
		case "1":
			res, err = day_02.Part1(inputPath)
		case "2":
			res, err = day_02.Part2(inputPath)
		default:
			fmt.Println("Invalid part. Use 1 or 2.")
			os.Exit(2)
		}
	case 3:
		switch part {
		case "1":
			res, err = day_03.Part1(inputPath)
		case "2":
			res, err = day_03.Part2(inputPath)
		default:
			fmt.Println("Invalid part. Use 1 or 2.")
			os.Exit(2)
		}
	case 4:
		switch part {
		case "1":
			res, err = day_04.Part1(inputPath)
		case "2":
			res, err = day_04.Part2(inputPath)
		default:
			fmt.Println("Invalid part. Use 1 or 2.")
			os.Exit(2)
		}
	case 8:
		switch part {
		case "1":
			res, err = day_08.Part1(inputPath)
		case "2":
			res, err = day_08.Part2(inputPath)
		default:
			fmt.Println("Invalid part. Use 1 or 2.")
			os.Exit(2)
		}
	case 10:
		switch part {
		case "1":
			res, err = day_10.Part1(inputPath)
		default:
			fmt.Println("Invalid part. Use 1 or 2.")
			os.Exit(2)
		}
	default:
		fmt.Printf("Day %d is not implemented yet.\n", dayNum)
		os.Exit(2)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)
}
