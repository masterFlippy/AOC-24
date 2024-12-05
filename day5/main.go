package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	content := string(file)

	stringSplit := strings.Split(content, "\n\n")
	ruleLines := strings.Split(stringSplit[0], "\n")
	pageNumbersLines := strings.Split(stringSplit[1], "\n")

	var rules []map[int]int

	for _, rule := range ruleLines {
		ruleSplit := strings.Split(rule, "|")
		if len(ruleSplit) == 2 {
			key, err1 := strconv.Atoi(ruleSplit[0])
			value, err2 := strconv.Atoi(ruleSplit[1])
			if err1 == nil && err2 == nil {
				ruleObj := map[int]int{key: value}
				rules = append(rules, ruleObj)
			}
		}
	}

	validArrays, invalidArrays := getArrays(pageNumbersLines, rules)
	validMiddleValues := getMiddleValues(validArrays)
	invalidMiddleValues := getMiddleValues(invalidArrays)

	validSum := 0
	invalidSum := 0
	for _, num := range validMiddleValues {
		validSum += num
	}
	for _, num := range invalidMiddleValues {
		invalidSum += num
	}

	fmt.Println("Part one: ", validSum)
	fmt.Println("Part two: ", invalidSum)
}

func getArrays(pageNumbersLines []string, rules []map[int]int) ([][]int, [][]int) {
	var validArrays [][]int
	var invalidArrays [][]int

	for _, pageNumbers := range pageNumbersLines {
		var pageArray []int
		for _, numStr := range strings.Split(pageNumbers, ",") {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}
			pageArray = append(pageArray, num)
		}
		seenValues := make(map[int]bool)
		invalid := false
		for _, pageNumber := range pageArray {
			for _, rule := range rules {
				for key, value := range rule {
					if key == pageNumber {
						if seenValues[value] {
							invalid = true
						}
					}
				}
			}

			seenValues[pageNumber] = true
		}
		if invalid {
			sortedPages := sortInvalidArray(pageArray, rules)
			invalidArrays = append(invalidArrays, sortedPages)
		} else {
			validArrays = append(validArrays, pageArray)
		}
	}

	return validArrays, invalidArrays
}

func sortInvalidArray(pageArray []int, rules []map[int]int) []int {
	sort.Slice(pageArray, func(i, j int) bool {
		for _, rule := range rules {
			for key, value := range rule {
				if pageArray[i] == key && pageArray[j] == value {
					return true
				}
				if pageArray[i] == value && pageArray[j] == key {
					return false
				}
			}
		}
		return pageArray[i] < pageArray[j]
	})

	return pageArray
}

func getMiddleValues(array [][]int) []int {
	var middleValues []int
	for _, array := range array {
		middleIndex := len(array) / 2
		middleValue := array[middleIndex]
		middleValues = append(middleValues, middleValue)
	}
	return middleValues
}
