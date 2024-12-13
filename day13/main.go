package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	X, Y int
}
type Machine struct {
	A, B, Prize Coordinate
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	machines, err := getMachines(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	partOne(machines)
	partTwo()

}

func partOne(machines []Machine) {
	maxInt := int(^uint(0) >> 1)
	tokens := 0
	for _, machine := range machines {
		cost := getMinCost(machine, maxInt)
		if cost != maxInt {
			tokens += cost
		}

	}
	fmt.Println("Part one: ", tokens)
}

func partTwo() {
	fmt.Println("Part two: ")
}

func getMachines(file *os.File) ([]Machine, error) {
	var machines []Machine
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Button A") {
			parts := strings.Split(line, ", ")
			xPart := strings.Split(parts[0], "+")[1]
			yPart := strings.Split(parts[1], "+")[1]
			xA, _ := strconv.Atoi(xPart)
			yA, _ := strconv.Atoi(yPart)

			scanner.Scan()
			line = scanner.Text()
			parts = strings.Split(line, ", ")
			xPart = strings.Split(parts[0], "+")[1]
			yPart = strings.Split(parts[1], "+")[1]
			xB, _ := strconv.Atoi(xPart)
			yB, _ := strconv.Atoi(yPart)

			scanner.Scan()
			line = scanner.Text()
			parts = strings.Split(line, ", ")
			xPart = strings.Split(parts[0], "=")[1]
			yPart = strings.Split(parts[1], "=")[1]
			xPrize, _ := strconv.Atoi(xPart)
			yPrize, _ := strconv.Atoi(yPart)

			machines = append(machines, Machine{
				A:     Coordinate{X: xA, Y: yA},
				B:     Coordinate{X: xB, Y: yB},
				Prize: Coordinate{X: xPrize, Y: yPrize},
			})
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return machines, nil
}

func getMinCost(machine Machine, maxInt int) int {
	xTarget, yTarget := machine.Prize.X, machine.Prize.Y
	costA, costB := 3, 1

	minCost := maxInt

	for a := 0; a <= xTarget/machine.A.X; a++ {
		for b := 0; b <= yTarget/machine.B.Y; b++ {
			xMove := a*machine.A.X + b*machine.B.X
			yMove := a*machine.A.Y + b*machine.B.Y

			if xMove == xTarget && yMove == yTarget {
				cost := costA*a + costB*b
				if cost < minCost {
					minCost = cost
				}
			}
		}
	}

	return minCost
}
