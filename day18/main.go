package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Byte struct {
	X, Y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	bytes, err := getBytes(file)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return
	}

	partOne(bytes)
	partTwo()

}

func partOne(bytes []Byte) {
	grid := getGrid(bytes, 71, 1024)
	rows, cols := len(grid), len(grid[0])
	start := Byte{0, 0}
	end := Byte{rows - 1, cols - 1}
	directions := []Byte{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}
	
	steps := findShortestPath(grid, start, end, rows, cols, directions)

	fmt.Println("Part one:", steps)

}

func partTwo() {

}

func getBytes(file *os.File) ([]Byte, error) {
	var bytes []Byte

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

		if len(parts) == 2 {
			x, errX := strconv.Atoi(parts[0])
			y, errY := strconv.Atoi(parts[1])
			if errX != nil || errY != nil {
				fmt.Println("Error parsing coordinates:", line)
				continue
			}

			bytes = append(bytes, Byte{X: x, Y: y})
		}
	}

	return bytes, nil
}

func getGrid(bytes []Byte, gridSize, byteCount int) [][]rune {
	grid := make([][]rune, gridSize)
	for i := 0; i < gridSize; i++ {
		grid[i] = make([]rune, gridSize)
		for j := 0; j < gridSize; j++ {
			grid[i][j] = '.'
		}
	}

	for i := 0; i < byteCount; i++ {
		b := bytes[i]
		if isInside(b.X, b.Y, gridSize, gridSize) {
			grid[b.Y][b.X] = '#'
		}
	}

	return grid
}

func findShortestPath(grid [][]rune, start, end Byte, rows, columns int, directions []Byte) int {
	if grid[start.X][start.X] == '#' || grid[end.X][end.X] == '#' {
		return -1
	}

	queue := list.New()
	queue.PushBack(Byte{0, 0})

	visited := make(map[Byte]bool)
	visited[start] = true

	steps := 0

	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			current := queue.Remove(queue.Front()).(Byte)

			if current == end {
				return steps
			}

			for _, d := range directions {
				neighbor := Byte{current.X + d.X, current.Y + d.Y}
				if isInside(neighbor.X, neighbor.Y, columns, rows) &&
					grid[neighbor.X][neighbor.Y] == '.' && !visited[neighbor] {
					queue.PushBack(neighbor)
					visited[neighbor] = true
				}
			}
		}
		steps++
	}

	return -1
}

func isInside(column, row, length, height int) bool {
	return column < length && row < height && column >= 0 && row >= 0
}
