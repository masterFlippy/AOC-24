package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	patterns, designs := getArrays(file)

	start := time.Now()
	partOne(patterns, designs)
	elapsed := time.Since(start)
	fmt.Printf("Part one %s\n", elapsed)

	start = time.Now()
	partTwo(patterns, designs)
	elapsed = time.Since(start)
	fmt.Printf("Part two %s\n", elapsed)

}

func partOne(patterns, designs []string) {
	count := 0
	cache := make(map[string]bool)
	for _, design := range designs {
		if canFormPattern(patterns, design, cache) {
			count++
		}
	}

	fmt.Println("Part one:", count)
}

func partTwo(patterns, designs []string) {
	cache := make(map[string]int)
	sum := 0
	for _, design := range designs {
		count := countPatterns(patterns, design, cache)
		sum += count
	}

	fmt.Println("Part Two:", sum)

}

func getArrays(file *os.File) ([]string, []string) {
	var patterns, designs []string

	scanner := bufio.NewScanner(file)
	isFirstArray := true

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			if isFirstArray {
				patterns = append(patterns, strings.Split(line, ", ")...)
				isFirstArray = false
			} else {
				designs = append(designs, line)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return patterns, designs
}

func canFormPattern(patterns []string, design string, cache map[string]bool) bool {
	if design == "" {
		return true
	}

	if result, exists := cache[design]; exists {
		return result
	}

	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) {
			if canFormPattern(patterns, design[len(pattern):], cache) {
				cache[design] = true
				return true
			}
		}
	}
	cache[design] = false

	return false
}

func countPatterns(patterns []string, design string, cache map[string]int) int {
	if design == "" {
		return 1
	}

	if result, exists := cache[design]; exists {
		return result
	}

	count := 0

	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) {
			count += countPatterns(patterns, design[len(pattern):], cache)
		}
	}

	cache[design] = count

	return count
}
