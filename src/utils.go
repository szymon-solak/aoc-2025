package src

import (
	"strconv"
	"strings"
)

func sliceAtoi(slice []string) []int {
	ints := make([]int, 0, len(slice))

	for _, item := range slice {
		asInt, err := strconv.Atoi(item)

		if err != nil {
			panic(err)
		}

		ints = append(ints, asInt)
	}

	return ints
}

type Cell struct {
	x   int
	y   int
	val byte
}

type Grid [][]Cell

func toGrid(input string) Grid {
	lines := strings.Split(input, "\n")

	grid := make(Grid, len(lines))

	for i, line := range lines {
		grid[i] = make([]Cell, len(line))
		for j, ch := range line {
			grid[i][j] = Cell{x: j, y: i, val: byte(ch)}
		}
	}

	return grid
}
