package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	for index, digit := range sortedArray {
		if(digit != ".") {
			digitInt, err := strconv.Atoi(digit)
			if err != nil {
				log.Fatalf("Failed to convert character to digit: %v", err)
			}
			sum += digitInt * index
		}
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

func getIdArray(digitArray []int) []string {
	var result []string

	for index, digit := range digitArray {
		if index%2 == 0 {
			temp := strings.Repeat(fmt.Sprintf("%d", index/2), digit)
			result = append(result, strings.Split(temp, "")...)
		} else {
			temp := strings.Repeat(".", digit)
			result = append(result, strings.Split(temp, "")...)
		}
	}

	return result
}

func sortArray(array []string) []string {
	i, j := 0, len(array)-1

	for i < j {
		if array[i] == "." {
			array[i], array[j] = array[j], array[i]
			j--
		} else {
			i++
		}
	}

	return array
}
