package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

type enginePart struct {
	Row   int
	Start int
	End   int
	Value int
}

type symbol struct {
	Row    int
	Column int
	isGear bool
}

func isSymbol(char rune) bool {
	return char != 46 && !unicode.IsDigit(char)
}

func parseInput(file *os.File) (*[]enginePart, *[]symbol) {
	var value, start, row int
	engineParts := []enginePart{}
	symbols := []symbol{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for column, char := range line {
			if unicode.IsDigit(char) {
				// Check if this is the first number
				if value == 0 {
					start = column
				}
				value = value*10 + int(char) - 48 // 48 is the unicode offset for ints
				continue
			}
			if value != 0 {
				// End of the number
				engineParts = append(engineParts, enginePart{
					Row:   row,
					Start: start,
					End:   column - 1,
					Value: value,
				})
				value = 0
			}
			if isSymbol(char) {
				symbols = append(symbols, symbol{Row: row, Column: column, isGear: char == 42})
			}
		}
		row++
	}
	return &engineParts, &symbols
}

func isNeighbour(part *enginePart, symbol *symbol) bool {
	// Is it within +/-1 rows
	rowDiff := symbol.Row - part.Row
	if rowDiff < 0 {
		rowDiff = rowDiff * -1
	}
	if rowDiff > 1 {
		return false
	}

	// is it within number columns +/-1
	isRightOfStart := symbol.Column-(part.Start-1) >= 0
	isLeftOfEnd := (part.End+1)-symbol.Column >= 0
	if isRightOfStart && isLeftOfEnd {
		return true
	}
	return false
}

func part1(engineParts *[]enginePart, symbols *[]symbol) int {
	total := 0
	seenPartIndexes := map[int]struct{}{}

	for _, symbol := range *symbols {
		for i, part := range *engineParts {
			if isNeighbour(&part, &symbol) {
				_, alreadySeen := seenPartIndexes[i]
				if !alreadySeen {
					total += part.Value
					seenPartIndexes[i] = struct{}{}
				}
			}
		}
	}
	return total
}

func part2(engineParts *[]enginePart, symbols *[]symbol) int {
	ratio := 0

	for _, symbol := range *symbols {
		symbolRatio := 0
		neighbours := 0
		for _, part := range *engineParts {
			if isNeighbour(&part, &symbol) {
				neighbours++
				if symbolRatio == 0 {
					symbolRatio = part.Value
				} else {
					symbolRatio *= part.Value
				}
			}
		}
		if neighbours == 2 {
			ratio += symbolRatio
		}
		neighbours = 0
		symbolRatio = 0
	}
	return ratio
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	engineParts, symbols := parseInput(file)
	fmt.Printf("Part 1: %d\n", part1(engineParts, symbols))
	fmt.Printf("Part 2: %d\n", part2(engineParts, symbols))
}
