package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type puzzle struct {
	Elf1     *elf
	Elf2     *elf
	Recipies []int
}

type elf struct {
	Score    int
	Position int
}

func main() {
	endAfter := parseInput()
	p := &puzzle{
		Elf1: &elf{
			Score:    3,
			Position: 0,
		},
		Elf2: &elf{
			Score:    7,
			Position: 1,
		},
		Recipies: []int{3, 7},
	}
	for i := 0; i < endAfter+11; i++ {
		p.doIteration()
	}
	fmt.Print("Answer: ")
	for _, i := range p.Recipies[endAfter : endAfter+10] {
		fmt.Print(i)
	}
	fmt.Println()

}

func (p *puzzle) doIteration() {
	n := numbersToAdd(p.Elf1.Score, p.Elf2.Score)
	p.Recipies = append(p.Recipies, n...)
	p.Elf1.Position += 1 + p.Elf1.Score
	p.Elf2.Position += 1 + p.Elf2.Score
	if p.Elf1.Position > len(p.Recipies)-1 {
		p.Elf1.Position = p.Elf1.Position % len(p.Recipies)
	}
	if p.Elf2.Position > len(p.Recipies)-1 {
		p.Elf2.Position = p.Elf2.Position % len(p.Recipies)
	}
	p.Elf1.Score = p.Recipies[p.Elf1.Position]
	p.Elf2.Score = p.Recipies[p.Elf2.Position]
}

func (p *puzzle) printStatus() {
	for index, val := range p.Recipies {
		switch index {
		case p.Elf1.Position:
			fmt.Print(fmt.Sprintf("(%d) ", val))
		case p.Elf2.Position:
			fmt.Print(fmt.Sprintf("[%d] ", val))
		default:
			fmt.Print(val, " ")
		}
	}
	fmt.Println()

}

func numbersToAdd(elf1 int, elf2 int) []int {
	nums := []int{}
	sum := elf1 + elf2
	sSum := strconv.Itoa(sum)
	for _, s := range sSum {
		i, err := strconv.Atoi(string(s))
		if err != nil {
			panic(err)
		}
		nums = append(nums, i)
	}
	return nums
}
func parseInput() int {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return i
	}
	return 0
}
