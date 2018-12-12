package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	startScore := int64(0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		startScore += parseLine(scanner.Text())
	}
	fmt.Println(startScore)
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
