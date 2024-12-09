package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	digitArray, err := getDigitArray(data)

	if err != nil {
		log.Fatal(err)
	}

	partOne(digitArray)
	partTwo()
}

func partOne(digitArray []int) {
	array := getIdArray(digitArray)
	sortedArray := sortArray(array)

	sum := 0
	for i, value := range sortedArray {
		sum += value * i
	}

	fmt.Println("Part one: ", sum)
}

func partTwo() {
	fmt.Println("Part two: ")
}

func getDigitArray(data []byte) ([]int, error) {
	numberString := string(data)
	var digits []int

	for _, char := range numberString {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			log.Fatalf("Failed to convert character to digit: %v", err)
		}
		digits = append(digits, digit)
	}

	return digits, nil
}

func getIdArray(digitArray []int) []int {
	var result []int

	value := 0
	for i, count := range digitArray {
		for j := 0; j < count; j++ {
			if i%2 == 0 {
				result = append(result, value)
			} else {
				result = append(result, -1)
			}
		}
		if i%2 == 0 {
			value++
		}
	}

	return result
}

func sortArray(array []int) []int {
	i, j := 0, len(array)-1

	for i < j {
		if array[i] == -1 {
			array[i], array[j] = array[j], array[i]
			j--
		} else {
			i++
		}
	}

	return trimArray(array)
}

func trimArray(array []int) []int {
	lastIndex := len(array) - 1
	for lastIndex >= 0 && array[lastIndex] == -1 {
		lastIndex--
	}
	return array[:lastIndex+1]
}
