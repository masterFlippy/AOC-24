package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"sort"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	mapList := parseInput(file)

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Println("Error seeking file:", err)
		return
	}

	pairs, err := parseInput2(file)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return
	}

	start := time.Now()
	partOne(mapList)
	elapsed := time.Since(start)
	fmt.Printf("Part one %s\n", elapsed)

	start = time.Now()
	partTwo(pairs)
	elapsed = time.Since(start)
	fmt.Printf("Part two %s\n", elapsed)
}

func partOne(mapList map[string][]string) {
	interConnectedSets := make(map[string][]string)
	for computer1 := range mapList {
		if computer1[0] != 't' {
			continue
		}
		for _, computer2 := range mapList[computer1] {
			for _, computer3 := range mapList[computer2] {
				if slices.Contains(mapList[computer3], computer1) {
					interConnectedSet := []string{computer1, computer2, computer3}
					slices.Sort(interConnectedSet)
					key := fmt.Sprintf("%s%s%s", interConnectedSet[0], interConnectedSet[1], interConnectedSet[2])
					interConnectedSets[key] = interConnectedSet
				}
			}
		}
	}

	fmt.Println("part one:", len(interConnectedSets))
}

func partTwo(pairs [][]string) {
	graph := make(map[string]map[string]struct{})
	for _, pair := range pairs {
		first, second := pair[0], pair[1]

		if _, exists := graph[first]; !exists {
			graph[first] = make(map[string]struct{})
		}
		if _, exists := graph[second]; !exists {
			graph[second] = make(map[string]struct{})
		}

		graph[first][second] = struct{}{}
		graph[second][first] = struct{}{}
	}

	var sets []map[string]struct{}

	for key, connections := range graph {
		added := false

		for _, set := range sets {
			containsAll := true
			for member := range set {
				if _, connected := connections[member]; !connected {
					containsAll = false
					break
				}
			}

			if containsAll {
				set[key] = struct{}{}
				added = true
				break
			}
		}

		if !added {
			newSet := make(map[string]struct{})
			newSet[key] = struct{}{}
			sets = append(sets, newSet)
		}
	}

	password := getPassword(sets)

	fmt.Println("part two:", password)
}

func parseInput(file *os.File) map[string][]string {
	scanner := bufio.NewScanner(file)
	mapList := make(map[string][]string)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "-")
		split1, split2 := split[0], split[1]
		mapList[split1] = append(mapList[split1], split2)
		mapList[split2] = append(mapList[split2], split1)
	}

	return mapList
}

func parseInput2(file *os.File) ([][]string, error) {
	var pairs [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		pairs = append(pairs, strings.Split(line, "-"))
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return pairs, nil
}

func getPassword(sets []map[string]struct{}) string {
	var largestSet map[string]struct{}
	maxSize := 0
	for _, set := range sets {
		if len(set) > maxSize {
			largestSet = set
			maxSize = len(set)
		}
	}

	var result []string
	for member := range largestSet {
		result = append(result, member)
	}
	sort.Strings(result)
	return strings.Join(result, ",")
}
