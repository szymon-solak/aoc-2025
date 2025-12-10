package src

import (
	"fmt"
	"strconv"
	"strings"
)

type MachineManual struct {
	targetIndicator     uint
	buttons             []uint
	joltageRequirements []uint
}

func parseIndicator(indicator string) uint {
	indicatorReplacer := strings.NewReplacer(
		"[", "",
		"]", "",
		".", "0",
		"#", "1",
	)

	num, err := strconv.ParseUint(
		indicatorReplacer.Replace(indicator),
		2,
		64,
	)

	if err != nil {
		panic(err)
	}

	return uint(num)
}

func parseButtons(buttons []string, indicators int) []uint {
	buttonReplacer := strings.NewReplacer("(", "", ")", "")
	var parsedButtons []uint

	for _, button := range buttons {
		var val uint

		for _, n := range sliceAtoi(strings.Split(buttonReplacer.Replace(button), ",")) {
			val = val | 0x1<<(indicators-n-1)
		}

		parsedButtons = append(parsedButtons, val)
	}

	return parsedButtons
}

func parseJoltage(joltage string) []uint {
	joltageReplacer := strings.NewReplacer("{", "", "}", "")
	var joltages []uint

	for _, j := range sliceAtoi(strings.Split(joltageReplacer.Replace(joltage), ",")) {
		joltages = append(joltages, uint(j))
	}

	return joltages
}

func parseManual(input string) []MachineManual {
	var manuals []MachineManual

	for line := range strings.SplitSeq(input, "\n") {
		parts := strings.Fields(line)

		manuals = append(manuals, MachineManual{
			targetIndicator:     parseIndicator(parts[0]),
			buttons:             parseButtons(parts[1:len(parts)-1], len(parts[0])-2),
			joltageRequirements: parseJoltage(parts[len(parts)-1]),
		})
	}

	return manuals
}

func getButtonCombinations(buttons []uint, presses int) [][]uint {
	combos := make([][]uint, 1)

	for i := 1; i <= presses; i++ {
		var next [][]uint

		for _, x := range combos {
			for _, y := range buttons {
				var t []uint
				t = append(t, x...)
				t = append(t, y)
				next = append(next, [][]uint{t}...)
			}
		}

		combos = next
	}

	return combos
}

func fewestPresses(manual MachineManual) int {
	presses := 1

	for {
		for _, combo := range getButtonCombinations(manual.buttons, presses) {
			var state uint

			for _, c := range combo {
				state = state ^ c
			}

			if state == manual.targetIndicator {
				return presses
			}
		}

		presses++
	}
}

func Day10Part1(input string) {
	result := 0

	for _, manual := range parseManual(input) {
		// fmt.Printf("[%b]\n", manual.targetIndicator)
		// fmt.Printf("[%b]\n\n", manual.buttons)
		// fmt.Println(fewestPresses(manual))
		result += fewestPresses(manual)
	}

	fmt.Println(result)

}

func Day10Part2(input string) {
}
