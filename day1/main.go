package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var leftArr []int
	var rightArr []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var left, right int
		_, err := fmt.Sscanf(scanner.Text(), "%d %d", &left, &right)
		if err != nil {
			log.Fatal(err)
		}

		leftArr = append(leftArr, left)
		rightArr = append(rightArr, right)
	}
	sort.Ints(leftArr)
	sort.Ints(rightArr)

	sum := 0
	sum2 := 0
	for i := 0; i < len(leftArr); i++ {
		if leftArr[i] > rightArr[i] {
			sum += leftArr[i] - rightArr[i]
		} else {
			sum += rightArr[i] - leftArr[i]
		}
		count := 0
		for j := 0; j < len(rightArr); j++ {
			if leftArr[i] == rightArr[j] {
				count++
			}
		}
		sum2 += count * leftArr[i]

	}
	fmt.Println(sum)
	fmt.Println(sum2)
}
