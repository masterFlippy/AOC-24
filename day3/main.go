package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
    file, err := os.ReadFile("input.txt")
    if err != nil {
        log.Fatalf("Error reading file: %v", err)
    }
    content := string(file)

    partOne(content)
    partTwo(content)
}

func partOne(content string)  {
	matches := getMatches(`mul\((\d{1,3}),(\d{1,3})\)`, content)
	sum := 0
	for _, match := range matches {
		sum += multiply(match[1], match[2])
	}

	fmt.Println("Part One: ", sum)
}

func partTwo(content string)  {
	matches := getMatches(`(?:mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\))`, content)
	sum := 0
	shouldDo := true
	for _, match := range matches {
		switch match[0] {
		case "don't()":
			shouldDo = false
		case "do()":
			shouldDo = true
		}
	
		if shouldDo {
			sum += multiply(match[1], match[2])
		}
	}
	
	fmt.Println("Part two: ", sum)
}

func getMatches(regex, content string) [][]string {
	re := regexp.MustCompile(regex)
	matches := re.FindAllStringSubmatch(content, -1)
	if len(matches) == 0 {
		fmt.Println("No matches found")
		return nil
	}
	return matches
}

func multiply(a, b string) int {
    num1, _ := strconv.Atoi(a)
    num2, _ := strconv.Atoi(b)
    return num1 * num2
}