package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type player struct {
	ID    int
	Score int
}

var endScore int

func main() {
	lastMarble, players := parseInput()
	players = playGame(players, lastMarble)
	highScore := 0
	for _, p := range players {
		if p.Score > highScore {
			highScore = p.Score
		}
	}
	fmt.Println("highscore is", highScore)
}

func playGame(players []*player, lastMarble int) []*player {
	currentIndex := 0
	marbles := make([]int, 0, 64976*10000)
	marbles = append(marbles, 0)
	p := 0
	marbleCounter := 1
	var done bool
mainLoop:
	for {
		currentIndex, done = players[p].generateNextIndex(marbles, currentIndex, marbleCounter, lastMarble)
		if done {
			break mainLoop
		}
		switch marbleCounter % 23 {
		case 0:
			players[p].Score += marbles[currentIndex]
			marbles = marbles[:currentIndex+copy(marbles[currentIndex:], marbles[currentIndex+1:])]
		default:
			marbles = append(marbles, 0)
			copy(marbles[currentIndex+1:], marbles[currentIndex:])
			marbles[currentIndex] = marbleCounter
			//printStatus(marbles, currentIndex, players[p])
		}
		p++
		if p >= len(players) {
			p = 0
		}
		marbleCounter++
		fmt.Println(marbleCounter)
	}
	return players
}

func printStatus(marbles []int, currentIndex int, p *player) {
	fmt.Printf("[%d] ", p.ID)
	for index, marble := range marbles {
		switch {
		case index == currentIndex:
			fmt.Printf("(%d) ", marble)
		default:
			fmt.Printf("%d ", marble)
		}
	}
	fmt.Printf("\n")
}

func (p *player) generateNextIndex(marbles []int, currentIndex int, marbleCounter int, lastMarble int) (int, bool) {
	var nextIndex int
	var done bool
	index1 := currentIndex + 1
	index2 := currentIndex + 2
	switch {
	case marbleCounter%23 == 0:
		nextIndex, done = p.handle23(marbles, currentIndex, marbleCounter, lastMarble)
	case index1 < len(marbles) && index2 < len(marbles):
		nextIndex = index2
	case index1 == len(marbles):
		nextIndex = 1
	case index2 == len(marbles):
		nextIndex = len(marbles)
	}
	return nextIndex, done

}

func (p *player) handle23(marbles []int, currentIndex int, marbleCounter int, lastMarble int) (int, bool) {
	var done bool
	p.Score += marbleCounter
	var removeIndex int
	switch {
	case currentIndex >= 7:
		removeIndex = currentIndex - 7
	case currentIndex < 7:
		removeIndex = len(marbles) - int(math.Abs(float64(currentIndex-7)))
	}
	if lastMarble <= marbleCounter+23 {
		done = true
	}
	return removeIndex, done

}

func parseInput() (int, []*player) {
	scanner := bufio.NewScanner(os.Stdin)
	var s []string
	for scanner.Scan() {
		s = strings.Split(scanner.Text(), " ")
	}
	parsedPlayers, err := strconv.Atoi(s[0])
	if err != nil {
		panic(err)
	}
	var players []*player
	for i := 1; i <= parsedPlayers; i++ {
		p := &player{
			ID: i,
		}
		players = append(players, p)
	}
	lastMarble, err := strconv.Atoi(s[6])
	if err != nil {
		panic(err)
	}
	lastMarble *= 100
	return lastMarble, players
}
