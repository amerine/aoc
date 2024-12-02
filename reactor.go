package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSafeReport(levels []int) bool {
	// Need at least 2 levels to compare
	if len(levels) < 2 {
		return false
	}

	// Track if we're increasing or decreasing
	increasing := true
	decreasing := true

	// Check each adjacent pair
	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i+1] - levels[i]
		
		// Check if difference is within allowed range (1-3)
		if abs(diff) < 1 || abs(diff) > 3 {
			return false
		}

		// Check if we maintain consistent direction
		if diff > 0 {
			decreasing = false
		} else if diff < 0 {
			increasing = false
		} else {
			// If diff is 0, it's not safe
			return false
		}
	}

	// Must be either all increasing or all decreasing
	return increasing || decreasing
}

func isSafeWithDampener(levels []int) bool {
	// First check if it's safe without removing any level
	if isSafeReport(levels) {
		return true
	}

	// Try removing each level one at a time
	for i := 0; i < len(levels); i++ {
		// Create a new slice without the current level
		dampened := make([]int, 0, len(levels)-1)
		dampened = append(dampened, levels[:i]...)
		dampened = append(dampened, levels[i+1:]...)

		// Check if removing this level makes it safe
		if isSafeReport(dampened) {
			return true
		}
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	safeCount := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Split line into numbers and convert to ints
		levelStrings := strings.Fields(line)
		levels := make([]int, len(levelStrings))

		for i, str := range levelStrings {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error converting string to number: %v\n", err)
				continue
			}
			levels[i] = num
		}

		if isSafeWithDampener(levels) {
			safeCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Number of safe reports: %d\n", safeCount)
}