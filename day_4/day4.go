package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filePath := "input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalPoints := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Split the line by ":"
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		// Split the second part by "|"
		winning, mine := strings.TrimSpace(parts[1]), ""
		if i := strings.Index(winning, "|"); i != -1 {
			winning, mine = winning[:i], winning[i+1:]
		}

		// Convert winning numbers to a set
		winningNumbers := make(map[string]struct{})
		for _, num := range strings.Fields(winning) {
			winningNumbers[num] = struct{}{}
		}

		// Count matching numbers
		matches := 0
		for _, num := range strings.Fields(mine) {
			if _, exists := winningNumbers[num]; exists {
				matches++
			}
		}

		// Calculate points
		if matches > 0 {
			totalPoints += 1 << uint(matches-1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Total Points:", totalPoints)
}
