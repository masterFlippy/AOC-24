package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Computer struct {
	registers          map[string]int
	program            []int
	instructionPointer int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	computer, err := parseInput(file)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return
	}

	partOne(computer)
	partTwo()

}

func partOne(computer Computer) {
	var outArr []int
	for computer.instructionPointer < len(computer.program) {
		instruction := computer.program[computer.instructionPointer]
		switch instruction {
		case 0:
			numerator := computer.registers["Register A"]
			denominator := int(math.Pow(2, float64(getComboOperand(computer, computer.program[computer.instructionPointer+1]))))
			computer.registers["Register A"] = numerator / denominator
		case 1:
			computer.registers["Register B"] ^= computer.program[computer.instructionPointer+1]
		case 2:
			computer.registers["Register B"] = getComboOperand(computer, computer.program[computer.instructionPointer+1]) % 8
		case 3:
			if computer.registers["Register A"] != 0 {
				computer.instructionPointer = computer.program[computer.instructionPointer+1]
				continue
			}
		case 4:
			computer.registers["Register B"] ^= computer.registers["Register C"]
		case 5:
			newValue := getComboOperand(computer, computer.program[computer.instructionPointer+1]) % 8
			outArr = append(outArr, newValue)

		case 6:
			numerator := computer.registers["Register A"]
			denominator := int(math.Pow(2, float64(getComboOperand(computer, computer.program[computer.instructionPointer+1]))))
			computer.registers["Register B"] = numerator / denominator

		case 7:
			numerator := computer.registers["Register A"]
			denominator := int(math.Pow(2, float64(getComboOperand(computer, computer.program[computer.instructionPointer+1]))))
			computer.registers["Register C"] = numerator / denominator

		}
		computer.instructionPointer += 2
	}

	var strBuilder strings.Builder
	for index, num := range outArr {
		strBuilder.WriteString(strconv.Itoa(num))
		if index < len(outArr)-1 {
			strBuilder.WriteString(", ")
		}
	}
	fmt.Println("Part one: ", strBuilder.String())
}

func partTwo() {

}

func parseInput(file *os.File) (Computer, error) {
	computer := Computer{
		registers: make(map[string]int),
		instructionPointer: 0,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "Register") {
			parts := strings.Split(line, ":")
			if len(parts) != 2 {
				return computer, fmt.Errorf("invalid register line: %s", line)
			}
			registerName := strings.TrimSpace(parts[0])
			value, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err != nil {
				return computer, fmt.Errorf("invalid register value: %s", line)
			}
			computer.registers[registerName] = value
		} else if strings.HasPrefix(line, "Program:") {
			parts := strings.Split(line, ":")
			if len(parts) != 2 {
				return computer, fmt.Errorf("invalid program line: %s", line)
			}
			values := strings.Split(strings.TrimSpace(parts[1]), ",")
			for _, v := range values {
				num, err := strconv.Atoi(strings.TrimSpace(v))
				if err != nil {
					return computer, fmt.Errorf("invalid program value: %s", v)
				}
				computer.program = append(computer.program, num)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return computer, err
	}

	return computer, nil
}

func getComboOperand(computer Computer, operand int) int {
	if operand == 0 || operand == 1 || operand == 2 || operand == 3 {
		return operand
	} else if operand == 4 {
		return computer.registers["Register A"]
	} else if operand == 5 {
		return computer.registers["Register B"]
	} else {
		return computer.registers["Register C"]
	}
}
