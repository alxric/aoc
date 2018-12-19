package main

import (
	"fmt"
	"math"
)

type puzzle struct {
	FuelCells   [][]fuelCell
	MaxPower    int
	MaxPowerPos string
}

type fuelCell struct {
	X          int
	Y          int
	PowerLevel int
}

func main() {
	p := generateGrid(300, 6548)
	findMaxPower(&p)
	fmt.Println(p.MaxPowerPos)
}

func findMaxPower(p *puzzle) {
	for y, row := range p.FuelCells {
		if y >= 298 {
			continue
		}
		for x, cell := range row {
			if x >= 298 {
				continue
			}
			totalPower := cell.PowerLevel
			totalPower += p.FuelCells[y][x+1].PowerLevel
			totalPower += p.FuelCells[y][x+2].PowerLevel
			totalPower += p.FuelCells[y+1][x].PowerLevel
			totalPower += p.FuelCells[y+1][x+1].PowerLevel
			totalPower += p.FuelCells[y+1][x+2].PowerLevel
			totalPower += p.FuelCells[y+2][x].PowerLevel
			totalPower += p.FuelCells[y+2][x+1].PowerLevel
			totalPower += p.FuelCells[y+2][x+2].PowerLevel
			if totalPower > p.MaxPower {
				p.MaxPower = totalPower
				p.MaxPowerPos = fmt.Sprintf("%d,%d", x+1, y+1)
			}
		}
	}
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
