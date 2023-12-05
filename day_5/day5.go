package main

import (
	//"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadInput(file string) string {
	b, err := os.ReadFile(file) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	// fmt.Println(b) // print the content as 'bytes'

	str := string(b) // convert content to a 'string'
	//fmt.Println("in the read: ",str)
	return str
	//fmt.Println(str) // print th
}

func GetSeeds(s string) []int {
	var seedsNumbers []int
	re := regexp.MustCompile(`seeds: (\d+ \d+ \d+ \d+)`)
	matchesSeeds := re.FindStringSubmatch(s)
	if len(matchesSeeds) > 0 {

		seedsNumbersInString := strings.Fields(matchesSeeds[1])
		ints := make([]int, len(seedsNumbersInString))
		for i, str := range seedsNumbersInString {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println(err)

			}
			ints[i] = num

		}
		seedsNumbers = ints
	}

	return seedsNumbers
}

func main() {
	input := ReadInput("sample.txt")
	fmt.Println(input)
	seedsNumbers := GetSeeds(input)
	fmt.Println("seed:", seedsNumbers)
	//mapRegex := regexp.MustCompile(`(\w+-to-\w+ map:.*?)\n\n`)
}
