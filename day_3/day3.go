package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"io"
)

func main() {
	// Read data from file
	filePath := "input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read content from file
	content, err := readFromFile(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Split lines
	lines := splitLines(content)

//	n := len(lines)
	m := len(lines[0])

	ans := 0

	for i, line := range lines {
		start := 0
		j := 0

		for j < m {
			start = j
			num := ""

			for j < m && isDigit(line[j]) {
				num += string(line[j])
				j++
			}

			if num == "" {
				j++
				continue
			}

			numVal := atoi(num)

			// Number ended, look around
			if isSymbol(i, start-1, lines) || isSymbol(i, j, lines) {
				ans += numVal
				continue
			}

			for k := start - 1; k <= j; k++ {
				if isSymbol(i-1, k, lines) || isSymbol(i+1, k, lines) {
					ans += numVal
					break
				}
			}
		}
	}

	fmt.Println(ans)
}

func readFromFile(file *os.File) (string, error) {
	var content strings.Builder
	_, err := io.Copy(&content, file)
	if err != nil {
		return "", err
	}
	return content.String(), nil
}

func splitLines(s string) []string {
	return strings.Split(strings.TrimSpace(s), "\n")
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func atoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

func isSymbol(i, j int, lines []string) bool {
	if !(0 <= i && i < len(lines) && 0 <= j && j < len(lines[0])) {
		return false
	}
	return lines[i][j] != '.' && !isDigit(lines[i][j])
}
