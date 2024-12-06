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

	partOne(copiedGrid)
	partTwo(grid)
}

func partOne(grid [][]rune) {
	directions := [][2]int{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}
	startX, startY, direction := findCursor(grid)

	x, y := startX, startY
	for {
		directionX, directionY := directions[direction][0], directions[direction][1]

		x, y = x+directionX, y+directionY

		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
			break
		}

		if grid[x][y] == '#' {
			direction = (direction + 1) % 4  
			x, y = x-directionX, y-directionY 
			continue
		}

		grid[x][y] = 'x'
	}

	fmt.Println("Part one: ", countGrid(grid, "x"))
}

func partTwo(grid [][]rune) {
	directions := [][2]int{
		{0, -1}, // up
		{1, 0},  // right
		{0, 1},  // down
		{-1, 0}, // left
	}
	rows, columns := len(grid), len(grid[0])
	startY, startX, startDirection := findCursor(grid)
	count := 0

	for column := 0; column < columns; column++ {
		for row := 0; row < rows; row++ {
			if grid[row][column] != '.' {
				continue
			}
			grid[row][column] = '#'
			directionX, directionY := directions[startDirection][0], directions[startDirection][1]
			x, y := startX, startY

			m := map[string]bool{}
			for isInside(x + directionX, y + directionY, columns, rows) {
				newX, newY := x + directionX, y + directionY
				if grid[newY][newX] == '#' {
					key := fmt.Sprintf("%d:%d:%d%d", x, y, directionX, directionY)
			
					if _, ok := m[key]; ok {
						count++
						break
					}
					m[key] = true
			
					directionX, directionY = -directionY, directionX
					continue
				}
				x, y = newX, newY
			}

			grid[row][column] = '.'
		}
	}

	fmt.Println("Part two: ", count)
}

func findCursor(grid [][]rune) (int, int, int) {
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {
			switch grid[row][column] {
			case '^':
				return row, column, 0 // up
			case '>':
				return row, column, 1 // right
			case 'v':
				return row, column, 2 // down
			case '<':
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
	return count + 1 // +1 for the starting position
}
func isInside(column, row, length, height int) bool {
	return column < length && row < height && column >= 0 && row >= 0
}
