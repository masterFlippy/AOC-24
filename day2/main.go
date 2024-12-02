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
	count := 0
	for _, array := range arrays {

		if isAscendingUsingSort(array) {
			if(checkLimit(array,1, 3)) {
				count++
			}
		} else if isDescendingUsingSort(array) {
			if(checkLimit(array,1, 3)) {
				count++
			}
		} 
	}
	fmt.Println(count)

}

func isAscendingUsingSort(array []int) bool {
	sorted := append([]int{}, array...) 
	sort.Ints(sorted)
	return reflect.DeepEqual(array, sorted)
}

func isDescendingUsingSort(array []int) bool {
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