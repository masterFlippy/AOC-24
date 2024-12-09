package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type IdBlock struct {
	Id  int
	Count int
}

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
	partTwo(digitArray)
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

func partTwo(digitArray []int) {
	idBlocks := getBlockArray(digitArray)
	sortedBlocks := sortBlockArray(idBlocks)

	index := 0
	sum := 0
	for len(sortedBlocks) > 0 {
		current := sortedBlocks[0]
		sortedBlocks = sortedBlocks[1:]

		if current.Id != -1 {
			sum += current.Id * index
			index++
			current.Count--
		} else {
			index += current.Count
			current.Count = 0
		}

		if current.Count > 0 {
			sortedBlocks = append([]IdBlock{current}, sortedBlocks...)
		}
	}

	fmt.Println("Part two: ", sum)
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

func getBlockArray(digitArray []int) []IdBlock {
	var idBLocks []IdBlock
	value := 0
	for i, count := range digitArray {
			if i%2 == 0 {
				idBLocks = append(idBLocks, IdBlock{Id: value, Count: count})
				value++
			} else {
				idBLocks = append(idBLocks, IdBlock{Id: -1, Count: count})
			}
	}

	return idBLocks
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

func sortBlockArray(array []IdBlock) []IdBlock {
	var sortedArray []IdBlock

	for i := 0; i < len(array); i++ {
		if array[i].Id == -1 {
			handleEmpty(array, &sortedArray, i)
		} else if array[i].Count > 0 {
			sortedArray = append(sortedArray, array[i])
		}
	}

	return sortedArray
}

func handleEmpty(array []IdBlock, sortedArray *[]IdBlock, i int) {
	for cursor := len(array) - 1; array[i].Count > 0 && cursor > i; cursor-- {
		if array[cursor].Id != -1 && array[cursor].Count <= array[i].Count {
			*sortedArray = append(*sortedArray, array[cursor])
			array[i].Count -= array[cursor].Count
			array[cursor].Id = -1
		}
	}

	if array[i].Count > 0 {
		*sortedArray = append(*sortedArray, array[i])
	}
}

func trimArray(array []int) []int {
	lastIndex := len(array) - 1
	for lastIndex >= 0 && array[lastIndex] == -1 {
		lastIndex--
	}
	return array[:lastIndex+1]
}
