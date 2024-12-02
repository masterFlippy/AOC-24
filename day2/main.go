package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var arrays [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		strNums := strings.Fields(line)

		var nums []int
		for _, str := range strNums {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}
			nums = append(nums, num)
		}

		arrays = append(arrays, nums)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	partOne(arrays)
	partTwo(arrays)

}

func partOne(arrays [][]int)  {
	count := 0
	for _, array := range arrays {
		if isSafe(array) {
			count++
		}
	}
	fmt.Println(count)
}

func partTwo(arrays [][]int)  {
	count := 0
	for _, array := range arrays {
		if isSafe(array) {
			count++
		} else {
			for i := range array {
				newArray:= append([]int(nil), array[:i]...) 
				newArray = append(newArray, array[i+1:]...)
				if isSafe(newArray) {
					count++
					break
				}
			}
		}	
	}
	fmt.Println(count)
}

func isSafe(array []int) bool {
	if isAscending(array) {
		if(checkLimit(array,1, 3)) {
			return true
		}
	} else if isDescending(array) {
		if(checkLimit(array,1, 3)) {
			return true
		}
	} 
	return false
}

func isAscending(array []int) bool {
	sorted := append([]int{}, array...) 
	sort.Ints(sorted)
	return reflect.DeepEqual(array, sorted)
}

func isDescending(array []int) bool {
	sorted := append([]int{}, array...)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
	return reflect.DeepEqual(array, sorted)
}

func checkLimit(array []int, minLimit int, maxLimit int) bool {
	for i := 0; i < len(array)-1; i++ {
		diff := math.Abs(float64(array[i] - array[i+1]))
		if diff < float64(minLimit) || diff > float64(maxLimit) {
			return false
		}
	}
	return true
}