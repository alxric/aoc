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
	parseInput()
}

type coordinate struct {
	ID int
	X  int
	Y  int
}

func parseInput() {
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
	drawGrid(xs, ys, sortXs[0], sortXs[len(sortXs)-1], sortYs[0], sortYs[len(sortYs)-1], coordinates)
}

func drawGrid(xs []int, ys []int, smallestX int, largestX int, smallestY int, largestY int, coordinates []coordinate) {
	var area int
	for y := smallestY; y <= largestY; y++ {
		for x := smallestX; x <= largestX; x++ {
			totalDiff := float64(0)
			for _, coord := range coordinates {
				totalDiff += calculateDiff(x, y, coord.X, coord.Y)
			}
			if totalDiff < 10000 {
				area++
			}

		}
	}
	fmt.Println(area)
}

func calculateDiff(a int, b int, c int, d int) float64 {
	return math.Abs(float64(a-c)) + math.Abs(float64(b-d))
}
