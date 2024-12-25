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

	keys, locks, height := parseInput(file)

	start := time.Now()
	partOne(keys, locks, height)
	elapsed := time.Since(start)
	fmt.Printf("Part one %s\n", elapsed)

}
func partOne(keys, locks [][]int, height int) {
	count := 0
	for _, key := range keys {
		for _, lock := range locks {
			if canFit(key, lock, height) {
				count++
			}
		}
	}
	fmt.Println("part one:", count)
}

func parseInput(file *os.File) ([][]int, [][]int, int) {
	var keys, locks [][]int
	var height int
	scanner := bufio.NewScanner(file)
	var grid []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if len(grid) > 0 {
				height = len(grid) - 1
				if strings.Contains(grid[0], "#") {
					locks = append(locks, getHeights(grid))
				} else if strings.Contains(grid[len(grid)-1], "#") {
					keys = append(keys, getHeights(grid))
				}
				grid = nil
			}
		} else {
			grid = append(grid, line)
		}
	}

	if len(grid) > 0 {
		if strings.Contains(grid[0], "#") {
			locks = append(locks, getHeights(grid))
		} else if strings.Contains(grid[len(grid)-1], "#") {
			keys = append(keys, getHeights(grid))
		}
	}

	return keys, locks, height
}

func getHeights(grid []string) []int {
	if len(grid) == 0 {
		return nil
	}

	heights := make([]int, len(grid[0]))

	for column := 0; column < len(grid[0]); column++ {
		height := -1
		for row := 0; row < len(grid); row++ {
			if grid[row][column] == '#' {
				height++
			}
		}
		heights[column] = height
	}

	return heights
}

func canFit(key, lock []int, height int) bool {
	for i := 0; i < len(key) && i < len(lock); i++ {
		if key[i]+lock[i] >= height {
			return false
		}
	}
	return true
}
