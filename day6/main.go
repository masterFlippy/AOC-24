package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid, err := getGrid(file)
	if err != nil {
		log.Fatal(err)
	}

	partOne(grid)
	partTwo(grid)
}

func partOne(grid [][]rune) {
	startX, startY, direction := findCursor(grid)
	directions := [][2]int{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}

	x, y := startX, startY
	for {
		directionX, directionY := directions[direction][0], directions[direction][1]

		x, y = x+directionX, y+directionY

		// check for hitting edge of grid
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
			break
		}

		// check if hit a guard
		if grid[x][y] == '#' {
			direction = (direction + 1) % 4   // rotate right
			x, y = x-directionX, y-directionY // revert the last move
			continue
		}

		// set as visited
		grid[x][y] = 'x'
	}

	fmt.Println("Part one: ", countGrid(grid, "x"))
}

func partTwo(grid [][]rune) {

	fmt.Println("Part two:")
}

func findCursor(grid [][]rune) (int, int, int) {
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {
			switch grid[row][column] {
			case '^':
				grid[row][column] = 'x' // change to visited
				return row, column, 0   // up
			case '>':
				grid[row][column] = 'x'
				return row, column, 1 // right
			case 'v':
				grid[row][column] = 'x'
				return row, column, 2 // down
			case '<':
				grid[row][column] = 'x'
				return row, column, 3 // left
			}
		}
	}
	panic("Cursor not found in grid")
}

func getGrid(file *os.File) ([][]rune, error) {
	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func countGrid(grid [][]rune, toCount string) int {
	count := 0
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {
			if grid[row][column] == rune(toCount[0]) {
				count++
			}
		}
	}
	return count
}
