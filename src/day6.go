package src

import (
	"fmt"
	"strconv"
	"strings"
)

type Problem struct {
	nums []int
	op   string
}

func parseProblemsLeftToRight(input string) []Problem {
	problems := []Problem{}
	lines := strings.Split(input, "\n")
	signs := strings.Fields(lines[len(lines)-1])

	fieldLines := [][]string{}

	for i := 0; i < len(lines)-1; i++ {
		fieldLines = append(fieldLines, strings.Fields(lines[i]))
	}

	for i := 0; i < len(fieldLines[0]); i++ {
		nums := []int{}

		for line := 0; line < len(fieldLines); line++ {
			asNum, err := strconv.Atoi(fieldLines[line][i])

			if err != nil {
				panic(err)
			}

			nums = append(nums, asNum)
		}

		problems = append(problems, Problem{nums: nums, op: strings.Trim(signs[i], " ")})
	}

	return problems
}

func transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := range xl {
		for j := range yl {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func parseProblemsRightToLeft(input string) []Problem {
	problems := []Problem{}
	lines := strings.Split(input, "\n")
	signs := strings.Fields(lines[len(lines)-1])

	numLines := [][]string{}
	for i := 0; i < len(lines)-1; i++ {
		numLines = append(numLines, strings.Split(lines[i], ""))
	}

	problemNums := make([][]int, 0, len(signs))
	problemNumIndex := 0

	for _, numArray := range transpose(numLines) {
		n := strings.Trim(strings.Join(numArray, ""), " ")

		if len(n) == 0 {
			problemNumIndex++
			continue
		}

		asNum, err := strconv.Atoi(n)

		if err != nil {
			panic(err)
		}

		if len(problemNums) <= problemNumIndex {
			next := []int{}
			problemNums = append(problemNums, next)
		}

		problemNums[problemNumIndex] = append(problemNums[problemNumIndex], asNum)
	}

	for index, numSet := range problemNums {
		problems = append(problems, Problem{nums: numSet, op: signs[index]})
	}

	return problems
}

func solveProblem(problem Problem) int {
	result := problem.nums[0]

	for i := 1; i < len(problem.nums); i++ {
		switch op := problem.op; op {
		case "+":
			{
				result += problem.nums[i]
			}

		case "*":
			{
				result *= problem.nums[i]
			}

		default:
			panic(fmt.Sprintf("Unhandled op: %s\n", op))
		}
	}

	return result
}

func Day6Part1(input string) {
	result := 0

	for _, problem := range parseProblemsLeftToRight(input) {
		result += solveProblem(problem)
	}

	fmt.Printf("Result = %d\n", result)
}

func Day6Part2(input string) {
	result := 0

	for _, problem := range parseProblemsRightToLeft(input) {
		result += solveProblem(problem)
	}

	fmt.Printf("Result = %d\n", result)
}
