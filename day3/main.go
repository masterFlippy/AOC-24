package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    buf := new(bytes.Buffer)
    buf.ReadFrom(file)
    content := buf.String()

    partOne(content)
    partTwo(content)

}

func partOne(content string)  {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(content, -1)

	if len(matches) == 0 {
		fmt.Println("No matches found")
		return
	}

	sum := 0
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		sum += num1 * num2
	}
	

	fmt.Println("Part One: ", sum)
}

func partTwo(content string)  {
	pattern := `(?:mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\))`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(content, -1)
	
	if len(matches) == 0 {
		fmt.Println("No matches found")
		return
	}

	sum := 0
	shouldDo := true
	for _, match := range matches {
		text := match[0]
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])

		switch text {
		case "don't()":
			shouldDo = false
		case "do()":
			shouldDo = true
		}
	
		if shouldDo {
			sum += num1 * num2
		}
	}
	

	fmt.Println("Part two: ", sum)

}