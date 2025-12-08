package day_08

import (
	"bufio"
	"os"
	"slices"
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
	jboxes := make([]JBox, 0)

	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		jboxes = append(jboxes, JBox{x, y, z})
	}

	distances := make([]Distance, 0)
	for i := 0; i < len(jboxes); i++ {
		for j := i + 1; j < len(jboxes); j++ {
			dist := calcDistance(jboxes[i], jboxes[j])
			distances = append(distances, Distance{jboxes[i], jboxes[j], dist})
		}
	}
	slices.SortFunc(distances, func(a, b Distance) int {
		return a.distance - b.distance
	})

	circuits := make([][]JBox, 0)
	mult := 0
	for len(distances) > 0 {
		minDist := distances[0]
		circuits = connectJBoxes(circuits, minDist.box1, minDist.box2)
		distances = distances[1:]
		if len(circuits) == 1 && len(circuits[0]) == len(jboxes) && mult == 0 {
			mult = minDist.box1.x * minDist.box2.x
		}
	}

	return mult, nil
}
