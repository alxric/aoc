package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	m := parseInput()
	biggestArea := 0
	for index, value := range m {
		if index != 200000 {
			if value > biggestArea {
				biggestArea = value
			}

		}
	}
	fmt.Println(biggestArea)
}

type coordinate struct {
	ID int
	X  int
	Y  int
}

func parseInput() map[int]int {
	var xs, ys []int
	var coordinates []coordinate
	scanner := bufio.NewScanner(os.Stdin)
	counter := 1
	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), ",")
		x, err := strconv.Atoi(strings.TrimSpace(vals[0]))
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(strings.TrimSpace(vals[1]))
		if err != nil {
			panic(err)
		}
		xs = append(xs, x)
		ys = append(ys, y)
		c := coordinate{
			ID: counter,
			X:  x,
			Y:  y,
		}
		coordinates = append(coordinates, c)
		counter++
	}
	sortXs := xs
	sortYs := ys
	sort.Ints(sortXs)
	sort.Ints(sortYs)
	return drawGrid(xs, ys, sortXs[0], sortXs[len(sortXs)-1], sortYs[0], sortYs[len(sortYs)-1], coordinates)
}

func drawGrid(xs []int, ys []int, smallestX int, largestX int, smallestY int, largestY int, coordinates []coordinate) map[int]int {
	m := make(map[int]int)
	var topLeftDiff, topRightDiff, bottomLeftDiff, bottomRightDiff float64
	var topLeft, topRight, bottomLeft, bottomRight coordinate
	for y := smallestY; y <= largestY; y++ {
		for x := smallestX; x <= largestX; x++ {
			smallestDiff := float64(2000000)
			var c coordinate
			for _, coord := range coordinates {
				diff := calculateDiff(x, y, coord.X, coord.Y)
				switch {
				case diff < smallestDiff:
					smallestDiff = diff
					c = coord
				case diff == smallestDiff:
					c = coordinate{
						ID: 200000,
					}
				}
				tlDiff := calculateDiff(smallestX, smallestY, coord.X, coord.Y)
				if tlDiff < topLeftDiff || topLeftDiff == 0 {
					topLeftDiff = diff
					topLeft = coord
				}
				trDiff := calculateDiff(largestX, smallestY, coord.X, coord.Y)
				if trDiff < topRightDiff || topRightDiff == 0 {
					topRightDiff = diff
					topRight = coord
				}
				blDiff := calculateDiff(smallestX, largestY, coord.X, coord.Y)
				if blDiff < bottomLeftDiff || bottomLeftDiff == 0 {
					bottomLeftDiff = diff
					bottomLeft = coord
				}
				brDiff := calculateDiff(largestX, largestY, coord.X, coord.Y)
				if brDiff < bottomRightDiff || bottomRightDiff == 0 {
					bottomRightDiff = diff
					bottomRight = coord
				}
			}
			m[c.ID]++

		}
	}
	m[topLeft.ID] = 0
	m[topRight.ID] = 0
	m[bottomRight.ID] = 0
	m[bottomLeft.ID] = 0
	return m
}

func calculateDiff(a int, b int, c int, d int) float64 {
	return math.Abs(float64(a-c)) + math.Abs(float64(b-d))
}
