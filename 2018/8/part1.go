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
	data, _ = parseTree(data, nextNodeID)
	totalMeta := 0
	for _, n := range nodes {
		for _, v := range n.MetaData {
			i, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			totalMeta += i
		}
	}
	fmt.Println(totalMeta)

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
