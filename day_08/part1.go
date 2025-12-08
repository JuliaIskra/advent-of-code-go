package day_08

import (
	"bufio"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type JBox struct {
	x, y, z int
}

type Distance struct {
	box1, box2 JBox
	distance   int
}

func Part1(filename string) (int, error) {
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

	var conns int
	if len(jboxes) == 20 {
		conns = 10
	} else {
		conns = 1000
	}

	circuits := make([][]JBox, 0)
	for i := 0; i < conns; i++ {
		minDist := distances[0]
		circuits = connectJBoxes(circuits, minDist.box1, minDist.box2)
		distances = distances[1:]
	}

	lengths := make([]int, 0)
	for _, circuit := range circuits {
		lengths = append(lengths, len(circuit))
	}
	sort.Ints(lengths)

	return lengths[len(lengths)-1] * lengths[len(lengths)-2] * lengths[len(lengths)-3], nil
}

func calcDistance(a JBox, b JBox) int {
	return (a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y) + (a.z-b.z)*(a.z-b.z)
}

func connectJBoxes(circuits [][]JBox, a, b JBox) [][]JBox {
	var foundA, foundB []JBox
	indA, indB := -1, -1

	for i, circuit := range circuits {
		if slices.Contains(circuit, a) {
			foundA = circuit
			indA = i
		}
		if slices.Contains(circuit, b) {
			foundB = circuit
			indB = i
		}
	}

	if indA == indB && indA != -1 {
		// both boxes are in the same circuit, nothing to do
		return circuits
	} else if indA == -1 && indB == -1 {
		// no circuit found yet, create a new one
		return append(circuits, []JBox{a, b})
	} else if indA != -1 && indB == -1 {
		// add b to circuit a
		circuits[indA] = append(circuits[indA], b)
		return circuits
	} else if indA == -1 && indB != -1 {
		// add a to circuit b
		circuits[indB] = append(circuits[indB], a)
		return circuits
	} else {
		// found two separate circuits, need to merge them
		merged := append(foundA, foundB...)
		var updated [][]JBox
		if indA < indB {
			updated = append(circuits[:indA], append(circuits[indA+1:indB], circuits[indB+1:]...)...)
		} else {
			updated = append(circuits[:indB], append(circuits[indB+1:indA], circuits[indA+1:]...)...)
		}
		return append(updated, merged)
	}
}
