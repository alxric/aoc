package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	startScore := int64(0)
	nums := make(map[int64]bool)
	nums[0] = true
	b, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	duplicateFound := false
	for !duplicateFound {
		for _, row := range strings.Split(strings.TrimSpace(string(b)), "\n") {
			startScore += parseLine(row)
			if _, ok := nums[startScore]; ok {
				fmt.Println("Duplicate found: ", startScore)
				duplicateFound = true
				break
			}
			nums[startScore] = true

		}
	}
}

func parseLine(line string) int64 {
	i, err := strconv.Atoi(line[1:])
	if err != nil {
		return 0
	}
	switch string(line[0]) {
	case "-":
		return 0 - int64(i)
	case "+":
		return int64(i)
	}
	return 0
}
