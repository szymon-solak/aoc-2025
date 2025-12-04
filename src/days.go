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
	}

	return days
}
