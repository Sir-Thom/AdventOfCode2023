package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isPossible(parts string) bool {
	lines := strings.Split(parts, "; ")

	for _, part := range lines {
		cubes := strings.Split(part, ", ")
		for _, cube := range cubes {
			numColor := strings.Split(cube, " ")
			num, _ := strconv.Atoi(numColor[0])
			color := numColor[1]

			switch color {
			case "red":
				if num > 12 {
					return false
				}
			case "green":
				if num > 13 {
					return false
				}
			case "blue":
				if num > 14 {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ans := 0

	for scanner.Scan() {
		line := scanner.Text()
		gameParts := strings.Split(line, ": ")
		game, parts := gameParts[0], gameParts[1]
		_id, _ := strconv.Atoi(strings.Split(game, " ")[1])

		if isPossible(parts) {
			ans += _id
		}
	}

	fmt.Println(ans)
}
