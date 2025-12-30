package day_02

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Part2(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ids := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		for _, str := range strings.Split(line, ",") {
			r := strings.Split(str, "-")
			start, _ := strconv.Atoi(r[0])
			end, _ := strconv.Atoi(r[1])
			ids = append(ids, []int{start, end})
		}
	}

	sum := 0
	for _, r := range ids {
		for id := r[0]; id <= r[1]; id++ {
			if hasDuplicatedSequence(id) {
				sum += id
			}
		}
	}

	return sum, nil
}

func hasDuplicatedSequence(id int) bool {
	idStr := strconv.Itoa(id)
	start := 1
	isRepeated := false
	for start < len(idStr) {
		repeat := idStr[0:start]
		rest := idStr[start:]
		repeated := strings.Repeat(repeat, len(rest)/len(repeat))
		if rest == repeated {
			isRepeated = true
			break
		} else {
			start = start + 1
			isRepeated = false
		}
	}
	return isRepeated
}
