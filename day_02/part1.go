package day_02

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
		sum2 := 0
		for id := r[0]; id <= r[1]; id++ {
			if isRepeatable(id) {
				sum2 += id
			}
		}
		sum += sum2
	}

	return sum, nil
}

func isRepeatable(id int) bool {
	idStr := strconv.Itoa(id)
	first, rest := idStr[0:len(idStr)/2], idStr[len(idStr)/2:]
	return rest == first
}
