package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)
type Coordinate struct {
	X, Y int
}

type Direction struct {
	directionX, directionY int
}

type Path struct {
	Positions []Coordinate
	Points int
}

type Paths struct {
	Paths []Path
}

type State struct {
	point     Coordinate
	direction string
	path      []Coordinate
	points    int
}


func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	grid := getGrid(file)

	directions := map[string]Direction{
		"north": {0, -1},
		"east":  {1, 0}, 
		"south": {0, 1},
		"west":  {-1, 0},
	}


	partOne(grid, directions)
	partTwo()

}

func partOne(grid [][]rune, directions map[string]Direction) {
	start := findPosition(grid, "S")
	end := findPosition(grid, "E")
	points := getPoints(grid, start, end, directions)
	fmt.Println(points)
}

func partTwo() {

}

func getGrid(file *os.File) [][]rune {
	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	return grid
}

func findPosition(grid [][]rune, focusCell string) Coordinate {
	focusRune := []rune(focusCell)[0]
	for i, row := range grid {
		for j, cell := range row {
			if cell == focusRune {
				return Coordinate{i, j}
			}
		}
	}
	return Coordinate{}
}

func getPoints(grid [][]rune, start, end Coordinate, directions map[string]Direction) int {
	directionNames := []string{"north", "east", "south", "west"}
	queue := []State{{start, "east", nil, -1000}}

	visited := make(map[string]bool)

	for len(queue) > 0 {
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].points < queue[j].points
		})

		current := queue[0]
		queue = queue[1:]

		if current.point == end {
			return current.points
		}

		key := fmt.Sprintf("%d,%d,%s", current.point.X, current.point.Y, current.direction)
		if visited[key] {
			continue
		}
		visited[key] = true

		nextPosition := Coordinate{
			current.point.X + directions[current.direction].directionX,
			current.point.Y + directions[current.direction].directionY,
		}

		if isInside(nextPosition, grid) {
			queue = append(queue, State{
				nextPosition,
				current.direction,
				nil,
				current.points + 1,
			})
		}
		currentIndex := indexOf(current.direction, directionNames)
		leftIndex := (currentIndex + len(directionNames) - 1) % len(directionNames)
		rightIndex := (currentIndex + 1) % len(directionNames)

		queue = append(queue,
			State{
				current.point,
				directionNames[leftIndex],
				nil,
				current.points + 1000,
			},
			State{
				current.point,
				directionNames[rightIndex],
				nil,
				current.points + 1000,
			},
		)
	}

	return -1
}

func indexOf(direction string, directionNames []string) int {
	for i, d := range directionNames {
		if d == direction {
			return i
		}
	}
	return -1
}

func isInside(point Coordinate, grid [][]rune) bool {
    return point.X >= 0 && point.X < len(grid) && point.Y >= 0 && point.Y < len(grid[0]) && grid[point.X][point.Y] != '#'
}