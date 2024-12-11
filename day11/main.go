package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	content := string(data)
	stones := strings.Fields(content)

	// partOneSlow(stones)
	partOne(stones)
	partTwo(stones)

}

func partOne(stones []string) {
    stoneMap := make(map[int]string)
    for i, stone := range stones {
        stoneMap[i] = stone
    }

    for i := 0; i < 25; i++ {
        stoneMap = blinkFast(stoneMap)
		fmt.Println(i, len(stoneMap))
    }
}

func partTwo(stones []string) {
    stoneMap := make(map[int]int)

    for _, stone := range stones {
		value, _ := strconv.Atoi(stone)
        stoneMap[value] =  stoneMap[value] + 1
    }

    for i := 0; i < 75; i++ {
        stoneMap = blinkFaster(stoneMap)
    }

	sum := 0
	for _, amount := range stoneMap {
		sum += amount
	}

	fmt.Println("Part two: ", sum)
}

func blinkSlow(stones []string) []string {
	for i := 0; i < len(stones); i++ {
		stone := stones[i]

		if stone == "0" {
			stones[i] = "1"
		} else if len(stone)%2 == 0 {
			left, right := stone[:len(stone)/2], stone[len(stone)/2:]
			tempRight := right

			if strings.Trim(right, "0") == "" {
				tempRight = "0" 
			} else {
				tempRight = strings.TrimLeft(right, "0")
			}

			stones = append(stones[:i], append([]string{left, tempRight}, stones[i+1:]...)...)
			i++ 
		}else {
			num, err := strconv.Atoi(stones[i])
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return nil
			}
			stones[i] = strconv.Itoa(num * 2024)
		}
	}

	return stones
}

func blinkFast(stoneMap map[int]string) map[int]string {
    newMap := make(map[int]string)
    newIndex := 0

    for _, stone := range stoneMap {
        if stone == "0" {
            newMap[newIndex] = "1"
            newIndex++
        } else if len(stone)%2 == 0 {
            left, right := stone[:len(stone)/2], stone[len(stone)/2:]
            tempRight := right

            if strings.Trim(right, "0") == "" {
                tempRight = "0"
            } else {
                tempRight = strings.TrimLeft(right, "0")
            }

            newMap[newIndex] = left
            newIndex++
            newMap[newIndex] = tempRight
            newIndex++
        } else {
            num, err := strconv.Atoi(stone)
            if err != nil {
                fmt.Println("Error converting string to int:", err)
                return nil
            }
            newMap[newIndex] = strconv.Itoa(num * 2024)
            newIndex++
        }
    }

    return newMap
}

func blinkFaster(stones map[int]int) map[int]int {
	newMap := make(map[int]int)
		for stone, amount := range stones {
			var newStones []int
			if stone == 0 {
				newStones = append(newStones, 1)
			} else if len(strconv.Itoa(stone))%2 == 0 {
				strStone := strconv.Itoa(stone)
				strLeft, strRight := strStone[:len(strStone)/2], strStone[len(strStone)/2:]
				left,_ := strconv.Atoi(strLeft)
				right,_ := strconv.Atoi(strRight)
				newStones = append(newStones, left)
				newStones = append(newStones, right)
			} else {
				newStones = append(newStones, stone*2024)
			}
			for _, newStone := range newStones {
				newMap[newStone] = newMap[newStone] + amount
			}
		}

		return newMap
}