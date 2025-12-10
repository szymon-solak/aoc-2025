package src

type Day struct {
	Part1 func(input string)
	Part2 func(input string)
}

func GetDays() map[int]Day {
	days := map[int]Day{
		1: {Part1: Day1Part1, Part2: Day1Part2},
		2: {Part1: Day2Part1, Part2: Day2Part2},
		3: {Part1: Day3Part1, Part2: Day3Part2},
		4: {Part1: Day4Part1, Part2: Day4Part2},
		5: {Part1: Day5Part1, Part2: Day5Part2},
		6: {Part1: Day6Part1, Part2: Day6Part2},
		7: {Part1: Day7Part1, Part2: Day7Part2},
		8: {Part1: Day8Part1, Part2: Day8Part2},
		9: {Part1: Day9Part1, Part2: Day9Part2},
		10: {Part1: Day10Part1, Part2: Day10Part2},
	}

	return days
}
