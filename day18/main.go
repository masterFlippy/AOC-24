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
	directions := []Byte{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}

	partOne(bytes, directions)
	partTwo(bytes, directions)

}

func partOne(bytes []Byte, directions []Byte) {
	grid := getGrid(bytes, 71, 1024)
	rows, cols := len(grid), len(grid[0])
	start := Byte{0, 0}
	end := Byte{rows - 1, cols - 1}

	steps := findShortestPath(grid, start, end, rows, cols, directions)

	fmt.Println("Part one:", steps)
}

func partTwo(bytes []Byte, directions []Byte) {
	gridSize := 71
	start := Byte{0, 0}
	end := Byte{70, 70}

	stopByte := getStopByte(bytes, gridSize, len(bytes), gridSize, gridSize, start, end, directions)

	fmt.Println("Part two:", stopByte)

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
		byte := bytes[i]
		if isInside(byte.X, byte.Y, gridSize, gridSize) {
			grid[byte.Y][byte.X] = '#'

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

			for _, direction := range directions {
				neighbor := Byte{current.X + direction.X, current.Y + direction.Y}
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

func getStopByte(bytes []Byte, gridSize, byteCount, rows, columns int, start, end Byte, directions []Byte) Byte {
	grid := make([][]rune, gridSize)
	for i := 0; i < gridSize; i++ {
		grid[i] = make([]rune, gridSize)
		for j := 0; j < gridSize; j++ {
			grid[i][j] = '.'
		}
	}

	for i := 0; i < byteCount; i++ {
		byte := bytes[i]
		if isInside(byte.X, byte.Y, gridSize, gridSize) {
			grid[byte.Y][byte.X] = '#'
			steps := findShortestPath(grid, start, end, rows, columns, directions)
			if steps == -1 {
				return byte
			}
		}
	}

	return Byte{0, 0}
}

func isInside(column, row, length, height int) bool {
	return column < length && row < height && column >= 0 && row >= 0
}
