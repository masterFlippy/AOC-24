package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Operation struct {
	Input1    string
	Input2    string
	Operation string
	Result    string
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	valueMap, operations := parseInput(file)

	start := time.Now()
	partOne(valueMap, operations)
	elapsed := time.Since(start)
	fmt.Printf("Part one %s\n", elapsed)

}
func partOne(valueMap map[string]int, operations []Operation) {
	remaining := true

	for remaining {
		remaining = false

		for _, operation := range operations {
			_, exists1 := valueMap[operation.Input1]
			_, exists2 := valueMap[operation.Input2]

			if exists1 && exists2 {
				switch operation.Operation {
				case "AND":
					if valueMap[operation.Input1] == 1 && valueMap[operation.Input2] == 1 {
						valueMap[operation.Result] = 1
					} else {
						valueMap[operation.Result] = 0
					}
				case "OR":
					if valueMap[operation.Input1] == 1 || valueMap[operation.Input2] == 1 {
						valueMap[operation.Result] = 1
					} else {
						valueMap[operation.Result] = 0
					}
				case "XOR":
					if (valueMap[operation.Input1] == 1 && valueMap[operation.Input2] == 0) ||
						(valueMap[operation.Input1] == 0 && valueMap[operation.Input2] == 1) {
						valueMap[operation.Result] = 1
					} else {
						valueMap[operation.Result] = 0
					}
				}
			} else {
				remaining = true
			}
		}
	}

	filteredMap := make(map[string]int)
	for key, value := range valueMap {
		if strings.HasPrefix(key, "z") {
			filteredMap[key] = value
		}
	}

	keys := make([]string, 0, len(valueMap))
	for key := range filteredMap {
		keys = append(keys, key)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	var binaryString string
	for _, key := range keys {
		binaryString += strconv.Itoa(filteredMap[key])
	}

	decimal, err := strconv.ParseInt(binaryString, 2, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Part one:", decimal)
}

func parseInput(file *os.File) (map[string]int, []Operation) {
	valueMap := make(map[string]int)
	operations := []Operation{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.Contains(line, "->") {
			parts := strings.Split(line, "->")
			if len(parts) != 2 {
				fmt.Println("Invalid operation line:", line)
				continue
			}

			operation := strings.TrimSpace(parts[0])
			result := strings.TrimSpace(parts[1])

			opParts := strings.Fields(operation)
			if len(opParts) != 3 {
				fmt.Println("Invalid operation format:", line)
				continue
			}

			operations = append(operations, Operation{
				Input1:    opParts[0],
				Operation: opParts[1],
				Input2:    opParts[2],
				Result:    result,
			})
		} else {
			parts := strings.Split(line, ":")
			if len(parts) != 2 {
				fmt.Println("Invalid value line:", line)
				continue
			}

			key := strings.TrimSpace(parts[0])
			value, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err != nil {
				fmt.Println("Invalid value for key", key, ":", parts[1])
				continue
			}

			valueMap[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return valueMap, operations
}
