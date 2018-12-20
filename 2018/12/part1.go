package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type puzzle struct {
	States       map[string]int
	CurrentState string
	Notes        []note
	Last100      int
	HundredsDiff int
	IndexShift   int
}

type note struct {
	Pattern string
	NextGen string
}

var finalGeneration = 50000000000

func main() {
	p := parseInput()
	p.States = make(map[string]int)
	for i := 1; i <= 1000; i++ {
		var shift int
		p.CurrentState, shift = genNextState(p.CurrentState, p.Notes)
		p.IndexShift += shift
		if i%100 == 0 || i == 20 {
			score := calculatePotScore(p.CurrentState, p.IndexShift)
			fmt.Println(fmt.Sprintf("#%d: %d, diff since last 100: %d", i, score, score-p.Last100))
			if score-p.Last100 > p.HundredsDiff {
				p.HundredsDiff = score - p.Last100
			}
			p.Last100 = score
		}
	}
	score := calculatePotScore(p.CurrentState, p.IndexShift)
	mFactor := finalGeneration / 100
	finalScore := mFactor*p.HundredsDiff + score - 10*p.HundredsDiff
	fmt.Println("Final score is", finalScore)
}

func calculatePotScore(s string, indexShift int) int {
	score := 0
	for index, c := range s {
		if string(c) == "#" {
			score += indexShift + index
		}
	}
	return score
}

func genNextState(currentState string, notes []note) (string, int) {
	var indexShift int
	currentState, indexShift = addDots(currentState)
	pots := make(map[int]bool)
	for _, n := range notes {
		startIndex := 0
	patternLoop:
		for {
			found := strings.Index(currentState[startIndex:], n.Pattern)
			switch {
			case found == -1:
				break patternLoop
			default:
				switch n.NextGen {
				case "#":
					pots[found+2+startIndex] = true
				case ".":
					pots[found+2+startIndex] = false
				}
			}
			startIndex += found + 1
		}
	}
	var nextState string
	for i := range currentState {
		if val, ok := pots[i]; ok {
			switch val {
			case true:
				nextState += "#"
			case false:
				nextState += "."
			}
		} else {
			nextState += "."
		}
	}
	return nextState, indexShift
}

func addDots(s string) (string, int) {
	indexShifted := 0
	for {
		if string(s[len(s)-3:]) == "..." {
			break
		}
		s += "."
	}
	for {
		if string(s[:3]) == "..." {
			break
		}
		s = "." + s
		indexShifted--
	}
	return s, indexShifted
}

func parseInput() *puzzle {
	p := &puzzle{}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	p.CurrentState = strings.Split(scanner.Text(), "initial state: ")[1]
	for scanner.Scan() {
		if scanner.Text() != "" {
			s := strings.Split(scanner.Text(), " ")
			n := note{
				Pattern: s[0],
				NextGen: s[2],
			}
			p.Notes = append(p.Notes, n)
		}
	}
	return p
}
