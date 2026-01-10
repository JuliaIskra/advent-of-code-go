package day_05

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start, end int
}

func Part1(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	freshIds := make([]Range, 0)
	availableCount := 0

	dbBreak := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			dbBreak = true
			continue
		}

		if !dbBreak {
			freshIds = append(freshIds, parseRange(line))
		} else {
			id, _ := strconv.Atoi(line)
			for _, r := range freshIds {
				if r.start <= id && id <= r.end {
					availableCount++
					break
				}
			}
		}
	}

	return availableCount, nil
}

func parseRange(line string) Range {
	ns := strings.Split(line, "-")
	start, _ := strconv.Atoi(ns[0])
	end, _ := strconv.Atoi(ns[1])
	return Range{start, end}
}
