package src

import (
	"fmt"
	"strconv"
	"strings"
)

func Day1Part1(input string) {
	dial := 50
	timesPointingAtZero := 0

	for element := range strings.SplitSeq(input, "\n") {
		if len(element) == 0 {
			continue
		}

		dir := element[0]
		moveBy, _ := strconv.Atoi(element[1:])

		if dir == 'R' {
			dial = (dial + moveBy) % 100
		}

		if dir == 'L' {
			dial = ((dial-moveBy)%100 + 100) % 100
		}

		if dial == 0 {
			timesPointingAtZero += 1
		}
	}

	fmt.Printf("Times at 0 = %d\n", timesPointingAtZero)
}

func Day1Part2(input string) {
	dial := 50
	timesPointingAtZero := 0

	for element := range strings.SplitSeq(input, "\n") {
		if len(element) == 0 {
			continue
		}

		dir := element[0]
		moveBy, _ := strconv.Atoi(element[1:])

		if dir == 'R' {
			target := (dial + moveBy) % 100
			timesPointingAtZero += (dial + moveBy) / 100
			dial = target
		}

		if dir == 'L' {
			target := ((dial-moveBy)%100 + 100) % 100

			if moveBy >= dial {
				timesPointingAtZero += min(1*dial, 1) + (moveBy-dial)/100
			}

			dial = target
		}

	}

	fmt.Printf("Times at 0 = %d\n", timesPointingAtZero)
}
