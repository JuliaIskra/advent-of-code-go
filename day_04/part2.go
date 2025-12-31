package day_04

import (
	"bufio"
	"os"
	"strings"
)

type Coord [2]int

func Part2(filename string) (int, error) {
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

	overall := 0
	removable := true
	for removable {
		accessibleRolls := make(map[Coord]struct{})

		// count and remember accessible rolls
		for row := 0; row < len(rolls); row++ {
			for col := 0; col < len(rolls[row]); col++ {
				if rolls[row][col] != "@" {
					continue
				}
				nCoords := getNeighbours(row, col)
				nRolls := make(map[Coord]struct{})
				for _, nCoord := range nCoords {
					nRow := nCoord[0]
					nCol := nCoord[1]
					if isInBounds(nRow, len(rolls)) && isInBounds(nCol, len(rolls[row])) && rolls[nRow][nCol] == "@" {
						nRolls[Coord{nRow, nCol}] = struct{}{}
					}
				}
				if len(nRolls) < 4 {
					accessibleRolls[Coord{row, col}] = struct{}{}
				}
			}
		}
		if len(accessibleRolls) == 0 {
			removable = false
		}
		overall += len(accessibleRolls)

		// remove accessible rolls
		for coord := range accessibleRolls {
			rolls[coord[0]][coord[1]] = "x"
		}
	}

	return overall, nil
}
