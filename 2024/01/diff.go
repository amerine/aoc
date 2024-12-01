package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"math"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		return
	}

	inputFile := os.Args[1]
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	var columnA, columnB []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) != 2 {
			fmt.Println("Invalid line format, skipping.")
			continue
		}

		a, errA := strconv.Atoi(fields[0])
		b, errB := strconv.Atoi(fields[1])
		if errA != nil || errB != nil {
			fmt.Printf("Error parsing numbers: %v, %v\n", errA, errB)
			continue
		}

		columnA = append(columnA, a)
		columnB = append(columnB, b)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	sort.Ints(columnA)
	sort.Ints(columnB)

	var distances []int
	sumDistances := 0
	for i := 0; i < len(columnA) && i < len(columnB); i++ {
		distance := int(math.Abs(float64(columnA[i] - columnB[i])))
		distances = append(distances, distance)
		sumDistances += distance
	}

	// Calculate similarity score
	similarityScore := 0
	for _, a := range columnA {
		count := 0
		for _, b := range columnB {
			if a == b {
				count++
			}
		}
		similarityScore += a * count
	}

	fmt.Println("Sorted Column A:", columnA)
	fmt.Println("Sorted Column B:", columnB)
	fmt.Println("Distances:", distances)
	fmt.Println("Sum of Distances:", sumDistances)
	fmt.Println("Similarity Score:", similarityScore)
}
