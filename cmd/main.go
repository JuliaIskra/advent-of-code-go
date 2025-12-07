package main

import (
	"advent-of-code-go/day_01"
	"fmt"
	"os"
)

func main() {
	//filename := os.Args[1]
	res, err := day_01.Part1("day_01/input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)

	res, err = day_01.Part2("day_01/input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)
}
