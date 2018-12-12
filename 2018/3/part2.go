package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	boxClaims, numClaims := parseInput()
	badClaims := make(map[int]bool)
	for _, row := range boxClaims {
		for _, claims := range row {
			if len(claims) > 1 {
				for _, claim := range claims {
					badClaims[claim.Elf] = true
				}
			}
		}
	}
	/*
		for claims, status := range badClaims {
			if !status {
				fmt.Println(claims, status)
			}
		}*/
	for i := 1; i <= numClaims; i++ {
		if !badClaims[i] {
			fmt.Println(i)
		}
	}
}

type claim struct {
	Elf    int
	Left   int
	Top    int
	Width  int
	Height int
}

func parseInput() (map[int]map[int][]*claim, int) {
	boxClaims := make(map[int]map[int][]*claim)
	totalClaims := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		totalClaims++
		c := parseCase(scanner.Text())
		for y := 0; y < c.Height; y++ {
			for x := 0; x < c.Width; x++ {
				if _, ok := boxClaims[c.Left+x]; !ok {
					boxClaims[c.Left+x] = make(map[int][]*claim)
				}

				boxClaims[c.Left+x][c.Top+y] = append(boxClaims[c.Left+x][c.Top+y], c)
			}
		}
	}
	return boxClaims, totalClaims
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
