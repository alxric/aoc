package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	parseInput()
}

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
	for scanner.Scan() {
		s := scanner.Text()
		ok := true
		for ok {
			s, ok = scanString(s)
		}
		fmt.Println(len(s))
	}
}
