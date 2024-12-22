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

func getSecretNumber(number int) int {
	number = (number*64 ^ number) % 16777216
	number = ((number/32 ^ number) % 16777216)
	number = ((number*2048 ^ number) % 16777216)
	return number
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
