package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	m := parseInput()
	finalOrder := ""
	for len(m) > 0 {
		var doNow []string
		for _, s := range m {
			if len(s.Dependencies) == 0 {
				doNow = append(doNow, s.Name)
			}
		}
		deletedLetter := workOrder(doNow, m)
		delete(m, deletedLetter)
		finalOrder += deletedLetter
	}
	fmt.Println(finalOrder)
}

func workOrder(letters []string, m map[string]*step) string {
	sort.Strings(letters)
	for _, s := range m {
		delete(s.Dependencies, letters[0])
	}
	return letters[0]
}

type step struct {
	Name         string
	Dependencies map[string]bool
}

func parseInput() map[string]*step {
	m := make(map[string]*step)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		if _, ok := m[s[1]]; !ok {
			m[s[1]] = &step{
				Name:         s[1],
				Dependencies: map[string]bool{},
			}
		}
		if ss, ok := m[s[7]]; !ok {
			m[s[7]] = &step{
				Name: s[7],
				Dependencies: map[string]bool{
					s[1]: true,
				},
			}
		} else {
			ss.Dependencies[s[1]] = true
		}
	}
	return m
}
