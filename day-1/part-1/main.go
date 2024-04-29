package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func processLine(line string) int {
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

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += processLine(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
