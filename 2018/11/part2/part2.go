package main

import (
	"fmt"
	"math"
)

type puzzle struct {
	FuelCells   [][]fuelCell
	MaxPower    int
	MaxPowerPos string
	MaxGridSize int
}

type fuelCell struct {
	X          int
	Y          int
	PowerLevel int
}

func main() {
	p := generateGrid(300, 6548)
	findMaxPower(&p)
	fmt.Println(fmt.Sprintf("%s,%d", p.MaxPowerPos, p.MaxGridSize))
}

func findMaxPower(p *puzzle) {
	p.MaxPower = -300
	for y, row := range p.FuelCells {
		for x := range row {
			for i := 0; i <= 300-int(math.Max(float64(x), float64(y))); i++ {
				totalPower := buildSquare(p, x, y, i)
				if totalPower > p.MaxPower {
					p.MaxPower = totalPower
					p.MaxPowerPos = fmt.Sprintf("%d,%d", x+1, y+1)
					p.MaxGridSize = i
				}
			}
		}
	}
}

func buildSquare(p *puzzle, x int, y int, i int) int {
	var nextRow, nextCell int
	power := 0
	for nextRow = 0; nextRow < i; nextRow++ {
		for nextCell = 0; nextCell < i; nextCell++ {
			power += p.FuelCells[y+nextRow][x+nextCell].PowerLevel
		}
	}
	return power
}

func generateGrid(gridSize int, serialNumber int) puzzle {
	p := puzzle{}
	for y := 1; y <= gridSize; y++ {
		var fuelCells []fuelCell
		for x := 1; x <= gridSize; x++ {
			f := fuelCell{
				X:          x,
				Y:          y,
				PowerLevel: calculatePowerLevel(x, y, serialNumber),
			}
			fuelCells = append(fuelCells, f)
		}
		p.FuelCells = append(p.FuelCells, fuelCells)
	}
	return p
}

func calculatePowerLevel(x int, y int, serial int) int {
	rackID := x + 10
	powerLevel := rackID * y
	powerLevel += serial
	powerLevel *= rackID
	powerLevel = powerLevel % int(math.Pow(10, float64(3)))
	powerLevel /= int(math.Pow(10, float64(3-1)))
	powerLevel -= 5
	return powerLevel
}
