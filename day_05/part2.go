package day_05

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	value    int
	position string
}

func Part2(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	freshIds := make([]Point, 0)
	freshCount := 0

	dbBreak := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			dbBreak = true

			slices.SortFunc(freshIds, func(a, b Point) int {
				if a.value != b.value {
					return a.value - b.value
				}
				return strings.Compare(a.position, b.position)
			})
			openCounter := 0
			mostLeft := 0
			for _, p := range freshIds {
				if p.position == "right" {
					openCounter--
				}
				if openCounter == 0 && p.position == "left" {
					mostLeft = p.value
				}
				if openCounter == 0 && p.position == "right" {
					freshCount = freshCount + p.value - mostLeft + 1
				}
				if p.position == "left" {
					openCounter++
				}
			}
		}

		if !dbBreak {
			freshIds = append(freshIds, parsePoint(line)...)
		} else {
			break
		}
	}

	return freshCount, nil
}

func parsePoint(line string) []Point {
	ns := strings.Split(line, "-")
	start, _ := strconv.Atoi(ns[0])
	end, _ := strconv.Atoi(ns[1])
	return []Point{Point{start, "left"}, Point{end, "right"}}
}
