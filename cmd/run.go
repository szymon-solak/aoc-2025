package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"aoc/src"

	"github.com/spf13/cobra"
)

func run(day int, part int, example bool) {
	inputFileNameSuffix := ""

	if example {
		inputFileNameSuffix = ".example"
	}

	inputFileName := fmt.Sprintf("input/day-%d%s.txt", day, inputFileNameSuffix)

	woringDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get working directory:", day)
		return
	}

	fmt.Println("Working directory:", woringDirectory)

	inputFilePath := filepath.Join(woringDirectory, inputFileName)

	fmt.Println("Using input:", inputFilePath)

	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		fmt.Println("Failed to read input file:", err)
		return
	}

	input := strings.TrimRight(string(data), "\n")

	fmt.Printf("Running solution for day %d\n", day)

	days := src.GetDays()

	if part == 1 {
		days[day].Part1(input)
		return
	}

	if part == 2 {
		days[day].Part2(input)
		return
	}

	days[day].Part1(input)
	days[day].Part2(input)
}

var runCmd = &cobra.Command{
	Use:     "run",
	Aliases: []string{"r"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		day, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid day:", err)
			return
		}

		part, err := cmd.Flags().GetInt("part")
		if err != nil {
			fmt.Println("Invalid part:", err)
			return
		}

		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			fmt.Println("Invalid example:", err)
			return
		}

		run(day, part, example)
	},
}

func init() {
	runCmd.Flags().Int("part", 0, "Only run a given part")
	runCmd.Flags().Bool("example", false, "Use example file")
	rootCmd.AddCommand(runCmd)
}
