package src

import "fmt"

type BeamRow map[int]int

func getStartingPoint(cells []Cell) BeamRow {
	row := BeamRow{}

	for _, cell := range cells {
		if cell.val == 'S' {
			row[cell.x] = 1
			return row
		}
	}

	panic("Starting point not found")
}

func getNextBeamRow(cells []Cell, previous BeamRow) (BeamRow, int) {
	row := BeamRow{}
	splitTimes := 0

	for _, cell := range cells {
		if cell.val == '^' && previous[cell.x] > 0 {
			row[cell.x-1] += previous[cell.x]
			row[cell.x+1] += previous[cell.x]
			splitTimes += 1
			continue
		}

		if previous[cell.x] > 0 {
			row[cell.x] += previous[cell.x]
		}
	}

	return row, splitTimes
}

func Day7Part1(input string) {
	grid := toGrid(input)
	lastRow := getStartingPoint(grid[0])
	split := 0

	for _, row := range grid[1:] {
		nextRow, splitTimes := getNextBeamRow(row, lastRow)
		lastRow = nextRow
		split += splitTimes
	}

	fmt.Printf("Result = %d\n", split)
}

func Day7Part2(input string) {
	grid := toGrid(input)
	lastRow := getStartingPoint(grid[0])

	for _, row := range grid[1:] {
		nextRow, _ := getNextBeamRow(row, lastRow)
		lastRow = nextRow
	}

	sum := 0

	for _, v := range lastRow {
		sum += v
	}

	fmt.Printf("Result = %d\n", sum)
}
