package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	numbers := parseInput(file)

	start := time.Now()
	partOne(numbers)
	elapsed := time.Since(start)
	fmt.Printf("Part one %s\n", elapsed)
	start = time.Now()
	partTwo(numbers)
	elapsed = time.Since(start)
	fmt.Printf("Part two %s\n", elapsed)

}

func partOne(numbers []int) {
	sum := 0
	for _, number := range numbers {
		for range 2000 {
			number = getSecretNumber(number)
		}
		sum += number
	}
	fmt.Println(sum)
}

func partTwo(numbers []int) {
	bananaMap := getBananaMap(numbers)

	result := 0
	for _, total := range bananaMap {
		if total > result {
			result = total
		}
	}

	fmt.Println(result)
}

func getSecretNumber(number int) int {
	number = (number*64 ^ number) % 16777216
	number = ((number/32 ^ number) % 16777216)
	number = ((number*2048 ^ number) % 16777216)
	return number
}

func getBananaMap(numbers []int) map[[4]int]int {
	bananaMap := map[[4]int]int{}

	for _, number := range numbers {
		seen := make(map[[4]int]bool)
		diffWindow := [4]int{10, 10, 10, 10}

		for range 2000 {
			before := number % 10
			number = getSecretNumber(number)
			current := number % 10

			diffWindow[0] = diffWindow[1]
			diffWindow[1] = diffWindow[2]
			diffWindow[2] = diffWindow[3]
			diffWindow[3] = current - before

			if !seen[diffWindow] {
				seen[diffWindow] = true
				bananaMap[diffWindow] += current
			}
		}
	}
	return bananaMap
}

func parseInput(file *os.File) []int {
	var numbers []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error converting line to int:", err)
			return nil
		}
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return numbers
}
