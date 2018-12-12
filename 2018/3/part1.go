package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	m := parseInput()
	total := 0
	for _, row := range m {
		for _, val := range row {
			if val >= 2 {
				total++
			}
		}
	}
	fmt.Println(total)
}

type claim struct {
	Elf    int
	Left   int
	Top    int
	Width  int
	Height int
}

func parseInput() map[int]map[int]int {
	m := make(map[int]map[int]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		c := parseCase(scanner.Text())
		for y := 0; y < c.Height; y++ {
			for x := 0; x < c.Width; x++ {
				if _, ok := m[c.Left+x]; !ok {
					m[c.Left+x] = make(map[int]int)
				}
				m[c.Left+x][c.Top+y]++
			}
		}
	}
	return m
}

func parseCase(s string) *claim {
	c := &claim{}
	vals := strings.Split(s, " ")
	if i, err := strconv.Atoi(vals[0][1:]); err == nil {
		c.Elf = i
	}
	pos := strings.Split(vals[2], ",")
	if i, err := strconv.Atoi(pos[0]); err == nil {
		c.Left = i
	}
	if i, err := strconv.Atoi(pos[1][0 : len(pos[1])-1]); err == nil {
		c.Top = i
	}
	size := strings.Split(vals[3], "x")
	if i, err := strconv.Atoi(size[0]); err == nil {
		c.Width = i
	}
	if i, err := strconv.Atoi(size[1]); err == nil {
		c.Height = i
	}
	return c
}
