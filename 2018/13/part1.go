package main

import (
	"bufio"
	"fmt"
	"os"
)

type puzzle struct {
	Carts  map[int]map[int]*cart
	Grid   [][]string
	CrashX int
	CrashY int
}

type cart struct {
	X              int
	Y              int
	Replaced       string
	Turns          int
	Direction      string
	MovedThisRound bool
}

func main() {
	p := parseInput()
	p.printStatus()
	for p.moveCarts() {
	}
	fmt.Println(fmt.Sprintf("%d,%d", p.CrashX, p.CrashY))
}

func (p *puzzle) moveCarts() bool {
	for y, row := range p.Grid {
		for x := range row {
			if c, ok := p.Carts[y][x]; ok {
				if c.MovedThisRound {
					continue
				}
				nextX := x
				nextY := y
				replaceWith := ""
				switch c.Direction {
				case "^":
					replaceWith = "|"
					nextY--
				case "v":
					replaceWith = "|"
					nextY++
				case ">":
					replaceWith = "-"
					nextX++
				case "<":
					replaceWith = "-"
					nextX--
				}
				if c.Replaced != "" {
					replaceWith = c.Replaced
					c.Replaced = ""
				}
				c.X = nextX
				c.Y = nextY
				c.MovedThisRound = true
				if _, ok := p.Carts[nextY]; !ok {
					p.Carts[nextY] = make(map[int]*cart)
				}
				if crash := p.Carts[nextY][nextX]; crash != nil {
					p.Grid[nextY][nextX] = "X"
					p.Grid[y][x] = replaceWith
					p.CrashX = nextX
					p.CrashY = nextY
					return false
				}
				p.Carts[nextY][nextX] = c
				delete(p.Carts[y], x)
				// Now we might need to change direction
				switch p.Grid[nextY][nextX] {
				case "+":
					//Intersection, check how many turns
					switch c.Turns {
					case 0:
						//What direction where we going?
						switch c.Direction {
						case "^":
							c.Direction = "<"
						case ">":
							c.Direction = "^"
						case "v":
							c.Direction = ">"
						case "<":
							c.Direction = "v"
						}
					case 2:
						switch c.Direction {
						case "^":
							c.Direction = ">"
						case ">":
							c.Direction = "v"
						case "v":
							c.Direction = "<"
						case "<":
							c.Direction = "^"
						}
					}
					c.Turns++
					c.Replaced = "+"
				case `\`:
					// Depends what direction we are going
					switch c.Direction {
					case "^":
						c.Direction = "<"
					case ">":
						c.Direction = "v"
					case "v":
						c.Direction = ">"
					case "<":
						c.Direction = "^"
					}
					c.Replaced = `\`
				case "/":
					// Depends what direction we are going
					switch c.Direction {
					case "^":
						c.Direction = ">"
					case ">":
						c.Direction = "^"
					case "v":
						c.Direction = "<"
					case "<":
						c.Direction = "v"
					}
					c.Replaced = "/"

				}
				if c.Turns == 3 {
					c.Turns = 0
				}
				p.Grid[nextY][nextX] = c.Direction
				p.Grid[y][x] = replaceWith
			}
		}
	}
	p.resetCartMoveStatus()
	return true
}

func (p *puzzle) checkCartStatus() {
	for _, row := range p.Carts {
		for _, c := range row {
			fmt.Println(c)
		}
	}
}

func (p *puzzle) resetCartMoveStatus() {
	for _, row := range p.Carts {
		for _, c := range row {
			c.MovedThisRound = false
		}
	}
}

func (p *puzzle) printStatus() {
	for _, row := range p.Grid {
		for _, c := range row {
			fmt.Print(c)
		}
		fmt.Println("")
	}

}

func parseInput() *puzzle {
	p := &puzzle{}
	p.Carts = make(map[int]map[int]*cart)
	scanner := bufio.NewScanner(os.Stdin)
	rowCounter := 0
	for scanner.Scan() {
		s := scanner.Text()
		var row []string
		for i, c := range s {
			switch string(c) {
			case "|":
			case "-":
			case "+":
			case "/":
			case `\`:
			case " ":
			default:
				ct := &cart{
					X:         i,
					Y:         rowCounter,
					Direction: string(c),
				}
				if _, ok := p.Carts[rowCounter]; !ok {
					p.Carts[rowCounter] = make(map[int]*cart)
				}
				p.Carts[rowCounter][i] = ct
			}
			row = append(row, string(c))
		}
		p.Grid = append(p.Grid, row)
		rowCounter++
	}
	return p
}
