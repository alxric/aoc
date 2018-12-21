package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type puzzle struct {
	Elf1                *elf
	Elf2                *elf
	Recipies            []int
	LookFor             []int
	CurrentlyLookingFor []int
	FoundAt             int
}

type elf struct {
	Score    int
	Position int
}

func main() {
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
		LookFor:  parseInput(),
	}
	p.CurrentlyLookingFor = p.LookFor
	for p.doIteration() {
	}
	fmt.Println(p.FoundAt)
}

func (p *puzzle) doIteration() bool {
	n := numbersToAdd(p.Elf1.Score, p.Elf2.Score)
	counter := 0
	for _, i := range n {
		if p.CurrentlyLookingFor[0] == i {
			if p.FoundAt == 0 {
				p.FoundAt = len(p.Recipies) + counter
			}
			if len(p.CurrentlyLookingFor) == 1 {
				return false
			}
			p.CurrentlyLookingFor = p.CurrentlyLookingFor[1:]
		} else {
			p.CurrentlyLookingFor = p.LookFor
			p.FoundAt = 0
			if p.CurrentlyLookingFor[0] == i {
				if p.FoundAt == 0 {
					p.FoundAt = len(p.Recipies) + counter
				}
				p.CurrentlyLookingFor = p.CurrentlyLookingFor[1:]
			}
		}
		counter++
	}
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
	return true
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
func parseInput() []int {
	var input []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		for _, c := range s {
			i, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			input = append(input, i)
		}
		return input
	}
	return nil
}
