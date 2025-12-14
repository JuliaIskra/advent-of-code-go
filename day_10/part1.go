package day_10

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Part1(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		configs := strings.Split(line, " ")

		var lights []bool
		var buttons [][]int
		for _, config := range configs {
			if config[0] == '[' {
				lights = make([]bool, len(config)-2)
				for i := 1; i < len(config)-1; i++ {
					lights[i-1] = config[i] == '#'
				}
			}
			if config[0] == '(' {
				idStr := strings.Split(config[1:len(config)-1], ",")
				idInt := make([]int, len(idStr))
				for i, id := range idStr {
					x, _ := strconv.Atoi(id)
					idInt[i] = x
				}
				buttons = append(buttons, idInt)
			}
		}

		getFastestSolution(lights, buttons)
	}

	return 0, nil
}

func getFastestSolution(desired []bool, buttons [][]int) int {
	//current := make([]bool, len(desired))
	return 0
}
