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
	word := "XMAS"
	reverseWord := "SAMX"
	rows := len(grid)
	columns := len(grid[0])
		count := 0
		directions := []struct{ drow, dlcolumn int }{
			{0, 1},  //  right
			{0, -1}, //  left
			{1, 0},  //  down
			{-1, 0}, //  up
			{1, 1},  //  down-right
			{-1, -1},//  up-left
			{1, -1}, //  down-left
			{-1, 1}, //  up-right
		}

		for row := 0; row < rows; row++ {
			for column := 0; column < columns; column++ {
				for _, direction := range directions {
					if checkWord(grid, row, column, direction.drow, direction.dlcolumn, rows, columns, word) {
						count++
					}
					if checkWord(grid, row, column, direction.drow, direction.dlcolumn,rows, columns, reverseWord) {
						count++
					}
				}
			}
		}

		fmt.Println("Part One: ", count / 2)
}

func partTwo(grid [][]rune)  {
    rows := len(grid)
    columns := len(grid[0])

    count := 0
	for row := 1; row < rows-1; row++ {
        for column := 1; column < columns-1; column++ {
            if grid[row][column] == 'A' { 
				topLeft := grid[row-1][column-1]
				topRight := grid[row-1][column+1]
				bottomLeft := grid[row+1][column-1]
				bottomRight := grid[row+1][column+1]

                if (topLeft == 'S' || topLeft == 'M') &&
                    (topRight == 'S' || topRight == 'M') &&
                    (bottomLeft == 'S' || bottomLeft == 'M') &&
                    (bottomRight == 'S' || bottomRight == 'M') &&
					(topLeft != bottomRight && topRight != bottomLeft) {
                    count++
                }
            }
        }
    }
	fmt.Println("Part two: ", count)
}

func getGrid(file *os.File) ([][]rune, error)  {
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

func checkWord(grid [][]rune, row, column, drow, dlcolumn, rows, columns int, word string) bool {
	for i := 0; i < len(word); i++ {
		row := row + i*drow
		column := column + i*dlcolumn
		if row < 0 || row >= rows || column < 0 || column >= columns || grid[row][column] != rune(word[i]) {
			return false
		}
	}
	return true
}