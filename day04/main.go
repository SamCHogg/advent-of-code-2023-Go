package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
)

var re = regexp.MustCompile(`\d+`)

func processCard(line string) int {
	game := strings.Split(line, ":")[1]
	halves := strings.Split(game, "|")
	winningNumbers, numbers := re.FindAllString(halves[0], -1), re.FindAllString(halves[1], -1)
	matches := 0

	for _, number := range numbers {
		if slices.Contains(winningNumbers, number) {
			matches++
		}
	}
	return matches
}

func part1(matches []int) int {
	sum := 0
	for _, card := range matches {
		if card == 0 {
			continue
		}
		sum += 1 << (card - 1)
	}
	return sum
}

func part2(matches []int) int {
	copies := make([]int, len(matches))
	sum := 0
	for i, match := range matches {
		copies[i]++
		sum += copies[i]
		for j := range match {
			copies[i+j+1] += copies[i]
		}
	}
	return sum
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	matches := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches = append(matches, processCard(line))
	}
	sumPart1 := part1(matches)
	sumPart2 := part2(matches)

	fmt.Printf("Part 1: %d\n", sumPart1)
	fmt.Printf("Part 2: %d\n", sumPart2)
}
