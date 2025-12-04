package src

import (
	"fmt"
	"slices"
	"strings"
)

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

func getNeighbouringCells(grid Grid, cell Cell) []Cell {
	neighbours := []Cell{}
	directions := []struct{ dx, dy int }{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	for _, dir := range directions {
		newX := cell.x + dir.dx
		newY := cell.y + dir.dy
		if newX >= 0 && newX < len(grid[0]) && newY >= 0 && newY < len(grid) {
			neighbours = append(neighbours, grid[newY][newX])
		}
	}

	return neighbours
}

func countIf(cells []Cell, fn func(cell Cell) bool) int {
	count := 0
	for _, cell := range cells {
		if fn(cell) {
			count++
		}
	}
	return count
}

func getRollsAccessibleByForklift(grid Grid) []Cell {
	rolls := []Cell{}

	for _, row := range grid {
		for _, cell := range row {
			if cell.val != '@' {
				continue
			}

			rollsAround := countIf(getNeighbouringCells(grid, cell), func(cell Cell) bool {
				return cell.val == '@'
			})

			if rollsAround < 4 {
				rolls = append(rolls, cell)
			}
		}
	}

	return rolls
}

func withoutRolls(grid Grid, rolls []Cell) Grid {
	for _, roll := range rolls {
		grid[roll.y][roll.x] = Cell{x: roll.x, y: roll.y, val: '.'}
	}
	return grid
}

func removeRolls(grid Grid) Grid {
	rollsToRemove := getRollsAccessibleByForklift(grid)

	if len(rollsToRemove) == 0 {
		return grid
	}

	return removeRolls(withoutRolls(grid, rollsToRemove))
}

func Day4Part1(input string) {
	fmt.Printf("Result = %d\n", len(getRollsAccessibleByForklift(toGrid(input))))
}

func Day4Part2(input string) {
	grid := toGrid(input)

	initialCount := countIf(slices.Concat(grid...), func(cell Cell) bool {
		return cell.val == '@'
	})

	finalCount := countIf(
		slices.Concat(removeRolls(grid)...),
		func(cell Cell) bool {
			return cell.val == '@'
		})

	fmt.Printf("Result = %d\n", initialCount-finalCount)
}
