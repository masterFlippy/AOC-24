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
	copiedGrid := make([][]rune, len(grid))
	for i := range grid {
		copiedGrid[i] = make([]rune, len(grid[i]))
		copy(copiedGrid[i], grid[i])
	}
	rows, columns := len(grid), len(grid[0])

	partOne(copiedGrid, rows, columns)
	partTwo(grid, rows, columns)
}

func partOne(grid [][]rune, rows, columns int) {
	result := make(map[struct {
		row, column int
	}]bool)

	for _, positions := range findAntennaPositions(grid, rows, columns) {
		for _, position := range combinations(positions) {
			rowOne, columnOne := position[0].row, position[0].column
			rowTwo, columnTwo := position[1].row, position[1].column

			positionOne := struct {
				row, column int
			}{2*rowTwo - rowOne, 2*columnTwo - columnOne}
			positionTwo := struct {
				row, column int
			}{2*rowOne - rowTwo, 2*columnOne - columnTwo}

			if isInside(positionOne.column, positionOne.row, columns, rows) {
				result[positionOne] = true
			}

			if isInside(positionTwo.column, positionTwo.row, columns, rows) {
				result[positionTwo] = true
			}
		}
	}
	fmt.Println("Part one: ", len(result))

}

func partTwo(grid [][]rune, rows, columns int) {
	result := make(map[struct {
		row, column int
	}]bool)

	for _, positions := range findAntennaPositions(grid, rows, columns) {
		for _, position := range combinations(positions) {
			rowOne, columnOne := position[0].row, position[0].column
			rowTwo, columnTwo := position[1].row, position[1].column

			directions := [][4]int{
				{rowTwo - rowOne, columnTwo - columnOne, rowOne, columnOne},
				{rowOne - rowTwo, columnOne - columnTwo, rowTwo, columnTwo},
			}

			for _, direction := range directions {
				changeRow, changeColumn, startRow, startColumn := direction[0], direction[1], direction[2], direction[3]
				for isInside(startColumn, startRow, columns, rows) {
					result[struct {
						row, column int
					}{startRow, startColumn}] = true
					startRow += changeRow
					startColumn += changeColumn
				}
			}
		}
	}
	fmt.Println("Part two: ", len(result))
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

func findAntennaPositions(grid [][]rune, rows, columns int) map[rune][]struct {
	row, column int
} {
	positions := make(map[rune][]struct {
		row, column int
	})

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			if grid[row][column] != '.' {
				positions[grid[row][column]] = append(positions[grid[row][column]], struct {
					row, column int
				}{row, column})
			}
		}
	}

	return positions
}

func isInside(column, row, length, height int) bool {
	return column < length && row < height && column >= 0 && row >= 0
}

func combinations(arr []struct {
	row, column int
}) [][2]struct {
	row, column int
} {
	var result [][2]struct {
		row, column int
	}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			result = append(result, [2]struct {
				row, column int
			}{arr[i], arr[j]})
		}
	}

	return result
}
