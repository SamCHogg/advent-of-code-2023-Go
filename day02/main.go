package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var cubeMap = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func splitCubes(round string) map[string]int {
	output := map[string]int{}
	cubes := strings.Split(round, ",")
	for _, cube := range cubes {
		parts := strings.Split(strings.TrimSpace(cube), " ")
		amount, _ := strconv.Atoi(parts[0])
		colour := parts[1]
		output[colour] = amount
	}
	return output
}

func splitRounds(line string) (int, []string) {
	parts := strings.Split(line, ":")
	gameID, _ := strconv.Atoi(parts[0][5:])

	rounds := strings.Split(parts[1], ";")
	return gameID, rounds
}

func part1(line string) int {
	gameID, rounds := splitRounds(line)
	for _, round := range rounds {
		cubes := splitCubes(round)
		for colour, amount := range cubes {
			maxCubes, ok := cubeMap[colour]
			if !ok {
				return 0
			}
			if amount > maxCubes {
				return 0
			}
		}
	}
	return gameID
}

func part2(line string) int {
	mostCubes := map[string]int{}

	_, rounds := splitRounds(line)
	for _, round := range rounds {
		cubes := splitCubes(round)
		for colour, amount := range cubes {
			if amount > mostCubes[colour] {
				mostCubes[colour] = amount
			}
		}
	}

	power := 0
	for _, amount := range mostCubes {
		if power == 0 {
			power = amount
		} else {
			power = power * amount
		}
	}
	return power
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

	fmt.Printf("Sum Part 1: %d\n", sumPart1)
	fmt.Printf("Sum Part 2: %d\n", sumPart2)
}
