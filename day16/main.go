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
	Points    int
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
	copiedGrid := make([][]rune, len(grid))
	for index := range grid {
		copiedGrid[index] = make([]rune, len(grid[index]))
		copy(copiedGrid[index], grid[index])
	}

	partOne(copiedGrid, directions)
	partTwo(grid, directions)

}

func partOne(grid [][]rune, directions map[string]Direction) {
	start, end := findPositions(grid, []string{"S", "E"})
	points := getPoints(grid, start, end, directions)
	fmt.Println("Part one: ", points)
}

func partTwo(grid [][]rune, directions map[string]Direction) {
	start, end := findPositions(grid, []string{"S", "E"})
	points := getPoints(grid, start, end, directions)
	paths := getPaths(grid, start, end, directions, points)
	uniquePositions := map[Coordinate]string{}

	for _, row := range paths {
		for _, coord := range row {
			uniquePositions[coord] = ""
		}
	}

	fmt.Println("Part two: ", len(uniquePositions))
}

func getGrid(file *os.File) [][]rune {
	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	return grid
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

		if isOK(nextPosition, grid) {
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

func getPaths(grid [][]rune, start, end Coordinate, directions map[string]Direction, target int) [][]Coordinate {
	var paths [][]Coordinate
	directionNames := []string{"north", "east", "south", "west"}
	queue := []State{{start, "east", []Coordinate{start}, -1000}}
	visited := make(map[string]int)
	
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.points > target {
			continue
		}

		key := fmt.Sprintf("%d,%d,%s", current.point.X, current.point.Y, current.direction)
		if points, exists := visited[key]; exists && points < current.points {
			continue
		}
		visited[key] = current.points

		if current.point == end && current.points == target {
			paths = append(paths, current.path)
			continue
		}

		nextPosition := Coordinate{
			current.point.X + directions[current.direction].directionX,
			current.point.Y + directions[current.direction].directionY,
		}

		if isOK(nextPosition, grid) {
			makePath := make([]Coordinate, len(current.path))
			copy(makePath, current.path)
			queue = append(queue, State{
				nextPosition,
				current.direction,
				append(makePath, nextPosition),
				current.points + 1,
			})
		}

		currentIndex := indexOf(current.direction, directionNames)
		leftIndex := (currentIndex + len(directionNames) - 1) % len(directionNames)
		rightIndex := (currentIndex + 1) % len(directionNames)

		for _, newDirection := range []string{directionNames[leftIndex], directionNames[rightIndex]} {
			queue = append(queue,
				State{
					current.point,
					newDirection,
					current.path,
					current.points + 1000,
				},
			)
		}
	}

	return paths
}


func indexOf(direction string, directionNames []string) int {
	for i, d := range directionNames {
		if d == direction {
			return i
		}
	}
	return -1
}

func isOK(point Coordinate, grid [][]rune) bool {
	return point.X >= 0 && point.X < len(grid) && point.Y >= 0 && point.Y < len(grid[0]) && grid[point.X][point.Y] != '#'
}