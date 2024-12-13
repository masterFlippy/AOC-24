package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
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
	start := time.Now()
	partOne(file, false)
	elapsed := time.Since(start)
	fmt.Printf("Part one %s\n", elapsed)
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Println("Error seeking file:", err)
		return
	}
	start = time.Now()
	partTwo(file, true)
	elapsed = time.Since(start)
	fmt.Printf("Part two %s\n", elapsed)
}

func partOne(file *os.File, isP2 bool) {
	machines, err := getMachines(file, isP2)
	if err != nil {
		fmt.Println(err)
		return
	}
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

func partTwo(file *os.File, isP2 bool) {
	machines, err := getMachines(file, isP2)
	if err != nil {
		fmt.Println(err)
		return
	}
	tokens := 0
	for _, machine := range machines {
		cost := getMinCostBigBrain(machine)
		tokens += cost

	}
	fmt.Println("Part two: ", tokens)
}

func getMachines(file *os.File, isP2 bool) ([]Machine, error) {
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

			if isP2 {
				xPrize += 10000000000000
				yPrize += 10000000000000
			}

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

func getMinCostBigBrain(machine Machine) int {
	aX, aY := float64(machine.A.X), float64(machine.A.Y)
	bX, bY := float64(machine.B.X), float64(machine.B.Y)
	prizeX, prizeY := float64(machine.Prize.X), float64(machine.Prize.Y)

	a := (prizeX - prizeY*bX/bY) / (aX - aY*bX/bY)
	b := (prizeY - a*aY) / bY

	roundA, roundB := math.Round(a), math.Round(b)

	if roundA >= 0 && roundB >= 0 {
		result1 := roundA*aX + roundB*bX
		result2 := roundA*aY + roundB*bY

		if result1 == prizeX && result2 == prizeY {
			return int(roundA)*3 + int(roundB)
		}
	}

	return 0
}
