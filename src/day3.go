package src

import (
	"fmt"
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

func parseBanks(bank string) []int {
	return sliceAtoi(strings.Split(bank, ""))
}

func maxInSlice(slice []int) (int, int) {
	pos := -1
	val := -1

	for index, value := range slice {
		if value > val {
			pos = index
			val = value
		}
	}

	return pos, val
}

func getBanksJoltage(bank []int, batteries int) []int {
	if len(bank) == 0 {
		return []int{0}
	}

	result := make([]int, 0, batteries)
	start := 0

	for len(result) < batteries {
		remaining := batteries - len(result)
		window := bank[start : len(bank)-remaining+1]

		pos, val := maxInSlice(window)

		result = append(result, val)
		start += pos + 1
	}

	return result
}

func listToValue(slice []int) int {
	var sb strings.Builder

	for i := range slice {
		sb.WriteString(fmt.Sprint(slice[i]))
	}

	asInt, err := strconv.Atoi(sb.String())

	if err != nil {
		panic(err)
	}

	return asInt
}

func Day3Part1(input string) {
	totalJoltage := 0

	for line := range strings.SplitSeq(input, "\n") {
		totalJoltage += listToValue(getBanksJoltage(parseBanks(line), 2))
	}

	fmt.Printf("Total Joltage = %d\n", totalJoltage)
}

func Day3Part2(input string) {
	totalJoltage := 0

	for line := range strings.SplitSeq(input, "\n") {
		totalJoltage += listToValue(getBanksJoltage(parseBanks(line), 12))
	}

	fmt.Printf("Total Joltage = %d\n", totalJoltage)
}
