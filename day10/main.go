package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	copiedGrid := make([][]int, len(grid))
	for i := range grid {
		copiedGrid[i] = make([]int, len(grid[i]))
		copy(copiedGrid[i], grid[i])
	}
	
	rows, columns := len(grid), len(grid[0])
	directions := [][]int{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
	}

	partOne(copiedGrid, directions, rows, columns)
	partTwo(grid, directions, rows, columns)
}

func partOne(grid, directions [][]int, rows, columns int) {
	sum := 0
	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			if grid[row][column] == 0 {
				score := getScore(grid, directions, row, column)
				sum += score
			}
		}
	}

	fmt.Println("Part one: ", sum)
}

func partTwo(grid, directions [][]int, rows, columns int) {
	sum := 0
	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			if grid[row][column] == 0 {
				rating := getRating(grid, directions, row, column)
				sum += rating
			}
		}
	}

	fmt.Println("Part two: ", sum)
}

func getGrid(file *os.File) ([][]int, error) {
	var grid [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, fmt.Errorf("invalid character '%c' in input: %v", char, err)
			}
			row[i] = num
		}
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func getScore(grid, directions [][]int, startRow, startColumn int) int {
	count := 0
	visited := make(map[[2]int]bool)
	stack := [][2]int{{startRow, startColumn}}

	maxRow, maxColumn := len(grid), len(grid[0])

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		currentValue := grid[current[0]][current[1]]

		if currentValue == 9 && !visited[current] {
			count++
			visited[current] = true
		}

		for _, direction := range directions {
			nextIndex := [2]int{current[0] + direction[0], current[1] + direction[1]}
			if isInside(nextIndex[1], nextIndex[0], maxColumn, maxRow) &&
				grid[nextIndex[0]][nextIndex[1]] == currentValue+1 && !visited[nextIndex] {
				stack = append(stack, nextIndex)
			}
		}
	}

	return count
}

func getRating(grid, directions [][]int, startRow, startColumn int) int {
	count := 0
	stack := [][2]int{{startRow, startColumn}}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		currentRow, currentColumn := current[0], current[1]
		currentValue := grid[currentRow][currentColumn]

		if currentValue == 9 {
			count++
			continue
		}

		for _, direction := range directions {
			nextRow, nextCol := currentRow+direction[0], currentColumn+direction[1]
			if isInside(nextCol, nextRow, len(grid[0]), len(grid)) &&
				grid[nextRow][nextCol] == currentValue+1 {
				stack = append(stack, [2]int{nextRow, nextCol})
			}
		}
	}

	return count
}

func isInside(column, row, length, height int) bool {
	return column < length && row < height && column >= 0 && row >= 0
}
