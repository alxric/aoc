package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	parseInput()
}

var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func scanString(s string) (string, bool) {
	for index, ss := range s {
		switch {
		case index == len(s)-1:
			return s, false
		case unicode.IsLower(ss):
			if unicode.ToUpper(ss) == rune(s[index+1]) {
				return s[:index] + s[index+2:], true
			}
		case unicode.IsUpper(ss):
			if unicode.ToLower(ss) == rune(s[index+1]) {
				return s[:index] + s[index+2:], true
			}
		}
	}
	return "", false
}

func parseInput() {
	scanner := bufio.NewScanner(os.Stdin)
	stringLengths := []int{}
	for scanner.Scan() {
		s := scanner.Text()
		var workingString string
		for _, letter := range alphabet {
			workingString = strings.Replace(s, letter, "", -1)
			workingString = strings.Replace(workingString, strings.ToUpper(letter), "", -1)
			ok := true
			for ok {
				workingString, ok = scanString(workingString)
			}
			stringLengths = append(stringLengths, len(workingString))
		}
	}
	sort.Ints(stringLengths)
	fmt.Println(stringLengths[0])
}
