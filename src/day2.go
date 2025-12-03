package src

import (
	"fmt"
	"strconv"
	"strings"
)

type IdRange struct {
	start uint64
	end   uint64
}

func asRanges(input string) []IdRange {
	rawRanges := strings.Split(input, ",")

	var ranges []IdRange
	for _, r := range rawRanges {
		bounds := strings.Split(r, "-")
		start := bounds[0]
		end := bounds[1]
		iStart, _ := strconv.ParseUint(start, 10, 64)
		iEnd, _ := strconv.ParseUint(end, 10, 64)

		ranges = append(ranges, IdRange{start: iStart, end: iEnd})
	}
	return ranges
}

func getRepeatedIdsInRange(prange IdRange) []uint64 {
	var invalidIds []uint64

	for id := prange.start; id <= prange.end; id++ {
		asString := fmt.Sprint(id)

		l := asString[:len(asString)/2]
		r := asString[len(asString)/2:]

		if l == r {
			invalidIds = append(invalidIds, id)
		}
	}

	return invalidIds
}

func getIdsWithRepeatedSequencesInRange(prange IdRange) []uint64 {
	var invalidIds []uint64

	for id := prange.start; id <= prange.end; id++ {
		asString := fmt.Sprint(id)
		chunkBy := len(asString) / 2

		for chunkBy > 0 {
			chunk := asString[0:chunkBy]

			for i := chunkBy; i <= len(asString)-chunkBy; i += chunkBy {
				nextChunk := asString[i : i+chunkBy]

				if chunk != nextChunk {
					break
				}

				if i == len(asString)-chunkBy {
					invalidIds = append(invalidIds, id)
					chunkBy = 0
					break
				}
			}

			chunkBy = chunkBy - 1
		}

	}

	return invalidIds
}

func Day2Part1(input string) {
	var sum uint64 = 0

	for _, v := range asRanges(input) {
		for _, id := range getRepeatedIdsInRange(v) {
			sum += id
		}
	}

	fmt.Printf("Sum = %d\n", sum)
}

func Day2Part2(input string) {
	var sum uint64 = 0

	for _, v := range asRanges(input) {
		for _, id := range getIdsWithRepeatedSequencesInRange(v) {
			sum += id
		}
	}

	fmt.Printf("Sum = %d\n", sum)
}
