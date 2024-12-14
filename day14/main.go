package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Coordinate struct {
	X, Y int
}
type Robot struct {
	Position Coordinate
	Velocity Coordinate
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	input := strings.TrimSpace(string(file))
	robots := getRobots(input)

	partOne(robots)
	partTwo(robots)

}

func partOne(robots []Robot) {
	positions := []Coordinate{}

	for _, robot := range robots {
		for i := 0; i < 100; i++ {
			robot.Position = moveRobot(robot, 101, 103)
		}

		positions = append(positions, robot.Position)
	}

	middleX, middleY := 50, 51
	quadrantCounts := [4]int{}

	for _, position := range positions {
		switch {
		case position.X < middleX && position.Y < middleY:
			quadrantCounts[0]++
		case position.X > middleX && position.Y < middleY:
			quadrantCounts[1]++
		case position.X < middleX && position.Y > middleY:
			quadrantCounts[2]++
		case position.X > middleX && position.Y > middleY:
			quadrantCounts[3]++
		}
	}

	fmt.Println("Part one: ", quadrantCounts[0]*quadrantCounts[1]*quadrantCounts[2]*quadrantCounts[3])
}

func partTwo(robots []Robot) {
	grid := [103][101]rune{}
	iteration := 1

	for {
		for row := 0; row < 103; row++ {
			for column := 0; column < 101; column++ {
				grid[row][column] = '.'
			}
		}

		for i := range robots {
			robots[i].Position = moveRobot(robots[i], 101, 103)
			grid[robots[i].Position.Y][robots[i].Position.X] = '#'
		}

		if isChristmasTree(grid) {
			fmt.Println(iteration)
			for row := 0; row < 103; row++ {
				for column := 0; column < 101; column++ {
					if grid[row][column] == '#' {
						fmt.Print("#")
					} else {
						fmt.Print(".")
					}
				}
				fmt.Println()
			}
			break
		}

		iteration++
	}
}

func getRobots(input string) []Robot {
	lines := strings.Split(input, "\n")
	robots := []Robot{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		var px, py, vx, vy int
		_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		if err != nil {
			fmt.Println("Error parsing line:", line, err)
			continue
		}

		robot := Robot{
			Position: Coordinate{X: px, Y: py},
			Velocity: Coordinate{X: vx, Y: vy},
		}
		robots = append(robots, robot)
	}

	return robots
}

func moveRobot(robot Robot, gridWidth, gridHeight int) Coordinate {
	robot.Position.X += robot.Velocity.X
	robot.Position.Y += robot.Velocity.Y

	if robot.Position.X < 0 {

		robot.Position.X = gridWidth + robot.Position.X
	} else if robot.Position.X >= gridWidth {
		robot.Position.X = robot.Position.X - gridWidth
	}

	if robot.Position.Y < 0 {
		robot.Position.Y = gridHeight + robot.Position.Y
	} else if robot.Position.Y >= gridHeight {
		robot.Position.Y = robot.Position.Y - gridHeight
	}
	return robot.Position
}

func isChristmasTree(grid [103][101]rune) bool {
	for row := 0; row < 103; row++ {
		for column := 0; column < 101; column++ {
			if matchesTreePattern(grid, column, row) {
				return true
			}
		}
	}
	return false
}

func matchesTreePattern(grid [103][101]rune, x, y int) bool {
	treePattern := [][]rune{
		{'.', '.', '#', '.', '.'},
		{'.', '#', '#', '#', '.'},
		{'#', '#', '#', '#', '#'},
	}

	for deltaY, row := range treePattern {
		for deltaX, cell := range row {
			newY, newX := y+deltaY, x+deltaX

			if (newY >= 103 || newX >= 101) || (grid[newY][newX] != cell) {
				return false
			}
		}
	}

	return true
}
