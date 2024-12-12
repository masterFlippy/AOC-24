package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	x, y int
}


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
	for index := range grid {
		copiedGrid[index] = make([]rune, len(grid[index]))
		copy(copiedGrid[index], grid[index])
	}
	rows := len(grid)
	columns := len(grid[0])

	partOne(copiedGrid, rows, columns)
	partTwo(grid)

}

func partOne(grid [][]rune, rows, columns int){
	sum := 0
	plots := calculatePerimeters(grid, rows, columns)
	for _, plot := range plots {
		fmt.Printf("Character %c, perimeter %direction, count %direction\n", plot.Char, plot.Perimeter, plot.Count)
		sum += plot.Perimeter * plot.Count
	}
	fmt.Println("Part one: ", sum)
}

func partTwo(grid [][]rune){
	fmt.Println("Part two: ")
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

func calculatePerimeters(grid [][]rune, rows, columns int) []struct {
	Char       rune
	Perimeter  int
	Count      int
} {
	visited := make(map[Point]bool)
	plot := []struct {
		Char      rune
		Perimeter int
		Count     int
	}{}

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			point := Point{row, column}
			if !visited[point] {
				char := grid[row][column]
				perimeter, count := calculatePerimeterAndCount(grid, point, visited, rows, columns)
				plot = append(plot, struct {
					Char      rune
					Perimeter int
					Count     int
				}{
					Char:      char,
					Perimeter: perimeter,
					Count:     count,
				})
			}
		}
	}

	return plot
}

func calculatePerimeterAndCount(grid [][]rune, start Point, visited map[Point]bool, rows, columns int) (int, int) {
	char := grid[start.x][start.y]

	directions := []Point{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}

	var perimeter, count int
	queue := []Point{start}
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		count++

		for _, direction := range directions {
			neighborX, neighborY := current.x+direction.x, current.y+direction.y
			neighbor := Point{current.x+direction.x, current.y+direction.y}

			if !isInside(neighborX, neighborY, rows, columns) || grid[neighborX][neighborY] != char{
				perimeter++
			} else if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
	}

	return perimeter, count
}


func isInside(column, row, length, height int) bool {
	return column < length && row < height && column >= 0 && row >= 0
}