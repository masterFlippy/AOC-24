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
    partTwo()

}

func partOne(content string)  {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(content, -1)

	if len(matches) == 0 {
		fmt.Println("No matches found")
		return
	}

	var tuples [][2]int
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		tuples = append(tuples, [2]int{num1, num2})
	}
	
	sum := 0
	for _, tuple := range tuples {
		sum += tuple[0] * tuple[1]
	}

	fmt.Println("Part One: ", sum)
}

func partTwo()  {
 fmt.Println("Part Two")
}