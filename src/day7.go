package src

import "fmt"

type BeamRow map[int]bool

func getStartingPoint(cells []Cell) BeamRow {
	row := BeamRow{}

	for _, cell := range cells {
		if cell.val == 'S' {
			row[cell.x] = true
		}
	}

	return row
}

func getNextBeamRow(cells []Cell, previous BeamRow) (BeamRow, int) {
	row := BeamRow{}
	splitTimes := 0

	for _, cell := range cells {
		if cell.val == '^' && previous[cell.x] {
			row[cell.x-1] = true
			row[cell.x+1] = true
			splitTimes += 1
			continue
		}

		if previous[cell.x] {
			row[cell.x] = true
		}
	}

	return row, splitTimes
}

func Day7Part1(input string) {
	grid := toGrid(input)
	prevRow := getStartingPoint(grid[0])
	split := 0

	for _, row := range grid[1:] {
		nextRow, splitTimes := getNextBeamRow(row, prevRow)
		prevRow = nextRow
		split += splitTimes
		fmt.Printf("splitTimes = %d, split = %d\n", splitTimes, split)
	}

	fmt.Printf("Result = %d\n", split)
}

func Day7Part2(input string) {
	// @todo: rewrite as graph

	grid := toGrid(input)
	prevRow := getStartingPoint(grid[0])

	for _, row := range grid[1:] {
		nextRow, _ := getNextBeamRow(row, prevRow)
		prevRow = nextRow
	}
}
