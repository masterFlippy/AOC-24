package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	calibrations, err := getArray(file)

	if err != nil {
		log.Fatal(err)
	}

	partOne(calibrations)
	partTwo(calibrations)
}

func partOne(calibrations [][]int) {
	operators := []string{"+", "*"}
	sum := 0
	for _, calibrationLine := range calibrations {
		if testCalibration(calibrationLine, operators) {
			sum += calibrationLine[0]
		}
	}

	fmt.Println("Part one: ", sum)
}

func partTwo(calibrations [][]int) {
	operators := []string{"+", "*", "||"}
	sum := 0
	for _, calibrationLine := range calibrations {
		if testCalibration(calibrationLine, operators) {
			sum += calibrationLine[0]
		}
	}
	fmt.Println("Part two: ", sum)
}

func getArray(input *os.File) ([][]int, error) {
	var result [][]int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		testValue, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			fmt.Println("Error parsing test value:", parts[0])
			continue
		}

		valuesString := strings.Fields(parts[1])
		var values []int
		for _, valuesString := range valuesString {
			value, err := strconv.Atoi(valuesString)
			if err != nil {
				fmt.Println("Error parsing value:", valuesString)
				continue
			}
			values = append(values, value)
		}

		combined := append([]int{testValue}, values...)
		result = append(result, combined)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return result, nil
}

func testCalibration(calibration []int, operators []string) bool {
	target := calibration[0]
	values := calibration[1:]
	gaps := len(values) - 1

	var operatorCombinations [][]string
	generateOperatorCombos(operators, gaps, []string{}, &operatorCombinations)
	for _, operatorCombination := range operatorCombinations {
		result := calculate(values, operatorCombination)

		if result == target {
			return true
		}
	}

	return false
}

func generateOperatorCombos(operators []string, gaps int, currentOperators []string, operatorCombinations *[][]string) {
	if len(currentOperators) == gaps {
		combination := make([]string, len(currentOperators))
		copy(combination, currentOperators)
		*operatorCombinations = append(*operatorCombinations, combination)
		return
	}

	for _, operator := range operators {
		currentOperators = append(currentOperators, operator)
		generateOperatorCombos(operators, gaps, currentOperators, operatorCombinations)
		currentOperators = currentOperators[:len(currentOperators)-1]
	}
}

func calculate(values []int, operators []string) int {
	result := values[0]
	for i := 0; i < len(operators); i++ {
		switch operators[i] {
		case "+":
			result += values[i+1]
		case "*":
			result *= values[i+1]
		case "||":
			result, _ = strconv.Atoi(strconv.Itoa(result) + strconv.Itoa(values[i+1]))
		}
	}
	return result
}
