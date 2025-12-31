package day_04

import (
	"bufio"
	"os"
	"strings"
)

func Part1(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rolls := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		rolls = append(rolls, strings.Split(line, ""))
	}

	accessibleCount := 0
	for row := 0; row < len(rolls); row++ {
		for col := 0; col < len(rolls[row]); col++ {
			if rolls[row][col] != "@" {
				continue
			}
			nCoords := getNeighbours(row, col)
			nRollsCount := 0
			for _, nCoord := range nCoords {
				nRow := nCoord[0]
				nCol := nCoord[1]
				if isInBounds(nRow, len(rolls)) && isInBounds(nCol, len(rolls[row])) && rolls[nRow][nCol] == "@" {
					nRollsCount++
				}
			}
			if nRollsCount < 4 {
				accessibleCount++
			}
		}
	}

	return accessibleCount, nil
}

func getNeighbours(row, col int) [][]int {
	return [][]int{
		{row - 1, col},
		{row - 1, col - 1},
		{row, col - 1},
		{row + 1, col - 1},
		{row + 1, col},
		{row + 1, col + 1},
		{row, col + 1},
		{row - 1, col + 1},
	}
}

func isInBounds(i, size int) bool {
	return 0 <= i && i < size
}
