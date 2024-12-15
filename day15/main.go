package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	grid, directions := getGridAndDirections(file)

	partOne(grid, directions)
	partTwo()

}

func partOne(grid [][]rune, directions []string) {
	var row, column, directionRow, directionColumn int
	for i, gridRow := range grid {
		for j, cell := range gridRow {
			if cell == '@' {
				row = i
				column = j
				break
			}
		}
	}

	total := 0
	for _, direction := range directions {
		directionRow, directionColumn = 0, 0
		switch direction {
		case "<":
			directionColumn = -1
		case ">":
			directionColumn = 1
		case "^":
			directionRow = -1
		case "v":
			directionRow = 1
		default:
			continue
		}

		if move(grid, row, column, directionRow, directionColumn) {
			row += directionRow
			column += directionColumn
		}
	}
	for i, gridRow := range grid {
		for j, cell := range gridRow {
			if cell == 'O' {
				total += i*100 + j
			}
		}
	}

	fmt.Println("Part one:", total)
}

func partTwo() {

}

func getGridAndDirections(file *os.File) ([][]rune, []string) {
	var grid [][]rune
	var directions []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		if strings.ContainsAny(line, "<>^v") {
			for _, char := range line {
				directions = append(directions, string(char))
			}
		} else {
			grid = append(grid, []rune(line))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, nil
	}

	return grid, directions
}

func move(grid [][]rune, row, column, directionRow, directionColumn int) bool {
	if grid[row][column] == '.' {
		return true
	}

	if grid[row][column] == '#' {
		return false
	}

	if move(grid, row+directionRow, column+directionColumn, directionRow, directionColumn) {
		cell := grid[row][column]
		grid[row+directionRow][column+directionColumn] = cell
		grid[row][column] = '.'
		return true
	}

	return false
}
