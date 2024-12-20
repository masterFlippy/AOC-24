package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)
type Coordinate struct {
	X, Y int
}

var directions = []Coordinate{
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
	{-1, 0}, // up
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	grid , err := getGrid(file)
	if err != nil {
		fmt.Println("Error getting grid:", err)
		return
	}
	start := time.Now()
	partOne(grid)
	elapsed := time.Since(start)
	fmt.Printf("Part one %s\n", elapsed)

}

func partOne(grid [][]rune) {
	start, end := findPositions(grid, []string{"S", "E"})
	count := getCheatCount(grid, start, end)

	fmt.Println("part one:", count)

}

func getCheatCount(grid [][]rune, start, end Coordinate) int {
	track := make(map[Coordinate]int)
	track[start] = 0
	current := start
	currentStep := 0

	for current != end {
		currentStep++
		x, y := current.X, current.Y
	
		for _, direction := range directions { 
			newX, newY := x+direction.X, y+direction.Y
			newCoordinate := Coordinate{X: newX, Y: newY}
	
			if isInside(grid, newCoordinate)  && !isInTrack(track, newCoordinate) {
				current = newCoordinate
				track[current] = currentStep
				break
			}
		}
	}
	
	count := 0
	for coordinate := range track {
		for _, direction := range directions {
			newX, newY := coordinate.X+direction.X, coordinate.Y+direction.Y
			newCoordinate := Coordinate{X: newX, Y: newY}
	
			if isInTrack(track, newCoordinate) {
				continue
			}
	
			neighborCoord := Coordinate{X: newX + direction.X, Y: newY + direction.Y}
			if isInTrack(track, neighborCoord) && track[neighborCoord]-track[coordinate] >= 102 {
				count++
			}
		}
	}
	return count
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

func findPositions(grid [][]rune, focusCells []string) (Coordinate, Coordinate) {
	var start, end Coordinate
	for _, focusCell := range focusCells {
	focusRune := []rune(focusCell)[0]
	for i, row := range grid {
		for j, cell := range row {
			if cell == focusRune {
				if focusCell == "S" {
					start = Coordinate{i, j}
				} else {
					end = Coordinate{i, j}
				}
			}
		}
	}
	}
	return start, end
}

func isInside(grid [][]rune, coordinate Coordinate) bool {
	rows, columns := len(grid), len(grid[0])
	return coordinate.X >= 0 && coordinate.Y >= 0 && coordinate.X < rows && coordinate.Y < columns && grid[coordinate.X][coordinate.Y] != '#'
}

func isInTrack(track map[Coordinate]int, position Coordinate) bool {
	_, exists := track[position]
	return exists
}