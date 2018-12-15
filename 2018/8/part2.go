package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	ID          int
	NumChildren int
	Children    []*node
	MetaData    []string
}

var nodes []*node

func main() {
	data := parseInput()
	nextNodeID := 1
	parseTree(data, nextNodeID)
	for i := len(nodes)/2 - 1; i >= 0; i-- {
		opp := len(nodes) - 1 - i
		nodes[i], nodes[opp] = nodes[opp], nodes[i]
	}

	score := countRootScore(nodes[0])
	fmt.Println("score is", score)
}

func countRootScore(n *node) int {
	score := 0
	if len(n.Children) == 0 {
		for _, v := range n.MetaData {
			i, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			score += i
		}
		return score
	}
	for _, v := range n.MetaData {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		if i <= len(n.Children) {
			score += countRootScore(n.Children[i-1])
		}
	}
	return score
}

func parseTree(data []string, nextNodeID int) ([]string, *node) {
	lenMetaData := genMetaData(data)
	numChildren, err := strconv.Atoi(data[0])
	if err != nil {
		panic(err)
	}
	n := &node{
		ID:          nextNodeID,
		NumChildren: numChildren,
	}
	switch {
	case n.NumChildren == 0:
		n.MetaData = data[2 : 2+lenMetaData]
		data = data[2+lenMetaData:]
		nodes = append(nodes, n)
		return data, n
	default:
		data = data[2:]
		for i := 0; i < n.NumChildren; i++ {
			nextNodeID++
			var child *node
			data, child = parseTree(data, nextNodeID)
			n.Children = append(n.Children, child)
		}
		n.MetaData = data[:lenMetaData]
		data = data[lenMetaData:]
	}
	nodes = append(nodes, n)
	return data, n
}

func genMetaData(data []string) int {
	numMetaData, err := strconv.Atoi(data[1])
	if err != nil {
		panic(err)
	}
	return numMetaData
}

func parseInput() (data []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data = strings.Split(scanner.Text(), " ")
	}
	return
}
