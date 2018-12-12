package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	parseInput()
}

func parseInput() {
	m := make(map[string]bool)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		m[scanner.Text()] = true
		countDiff(scanner.Text(), m)
	}

}

func countDiff(s string, m map[string]bool) {
	for entry := range m {
		diff := 0
		diffIndex := 0
		for index, char := range s {
			if string(char) != string(entry[index]) {
				diff++
				diffIndex = index
			}
		}
		if diff < 2 && diff > 0 {
			result := s[:diffIndex] + s[diffIndex+1:]
			fmt.Println(result)
		}
	}
}
