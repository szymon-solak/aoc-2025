package src

import (
	"strconv"
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
