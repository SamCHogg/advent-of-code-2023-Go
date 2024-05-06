package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var numbers = map[string]string{
	"one":   "one1one",
	"two":   "two2two",
	"three": "three3three",
	"four":  "four4four",
	"five":  "five5five",
	"six":   "six6six",
	"seven": "seven7seven",
	"eight": "eight8eight",
	"nine":  "nine9nine",
}

func part1(line string) int {
	firstPos := strings.IndexAny(line, "0123456789")
	lastPos := strings.LastIndexAny(line, "0123456789")
	if firstPos == -1 {
		return 0
	}

	first := string(line[firstPos])
	last := string(line[lastPos])

	output, err := strconv.Atoi(first + last)
	if err != nil {
		return 0
	}
	return output
}

func part2(line string) int {
	for word, digit := range numbers {
		line = strings.ReplaceAll(line, word, digit)
	}
	return part1(line)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var sumPart1, sumPart2 int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sumPart1 += part1(line)
		sumPart2 += part2(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sum of Part 1: %d\n", sumPart1)
	fmt.Printf("Sum of Part 2: %d\n", sumPart2)
}
