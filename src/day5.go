package src

import (
	"fmt"
	"strconv"
	"strings"
)

type Range struct {
	from int
	to   int
}

func parseRange(r string) Range {
	parts := strings.Split(r, "-")

	from, err := strconv.Atoi(parts[0])

	if err != nil {
		panic(err)
	}

	to, err := strconv.Atoi(parts[1])

	if err != nil {
		panic(err)
	}

	return Range{from: from, to: to}
}

func parseInput(input string) ([]Range, []int) {
	parts := strings.SplitN(input, "\n\n", 2)

	rangeLines := strings.Split(parts[0], "\n")

	ranges := []Range{}

	for _, line := range rangeLines {
		ranges = append(ranges, parseRange(line))
	}

	ids := sliceAtoi(strings.Split(parts[1], "\n"))

	return ranges, ids
}

func getFreshIds(freshRanges []Range, ids []int) []int {
	freshIds := []int{}

	for _, id := range ids {
		for _, r := range freshRanges {
			if id >= r.from && id <= r.to {
				freshIds = append(freshIds, id)
				break
			}
		}
	}

	return freshIds
}

func combineRanges(ranges []Range) []Range {
	combinedRanges := []Range{}

OUTER:
	for _, r := range ranges {
		for index, cr := range combinedRanges {
			if r.from <= cr.to && r.to >= cr.from {
				combinedRanges[index] = Range{from: min(r.from, cr.from), to: max(r.to, cr.to)}
				continue OUTER
			}
		}

		combinedRanges = append(combinedRanges, r)
	}

	if len(ranges) == len(combinedRanges) {
		return ranges
	}

	return combineRanges(combinedRanges)
}

func countRangeSum(ranges []Range) int {
	count := 0

	for _, r := range ranges {
		count += (r.to - r.from + 1)
	}

	return count
}

func Day5Part1(input string) {
	fmt.Printf("Fresh Ids = %d\n", len(getFreshIds(parseInput(input))))
}

func Day5Part2(input string) {
	ranges, _ := parseInput(input)
	fmt.Printf("All Fresh Ids = %d\n", countRangeSum(combineRanges(ranges)))
}
