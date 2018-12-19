package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"
)

type puzzle struct {
	Lights    []*light
	SmallestX int
	LargestX  int
	SmallestY int
	LargestY  int
	LastY     int
	YsInARow  int
}

type light struct {
	X    int
	Y    int
	XVel int
	YVel int
}

func main() {
	var data []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	p := parseInput(data)

	for i := 0; i < 50000; i++ {
		ysinarow := 0
		ys := make(map[int]int)
		for _, light := range p.Lights {
			ys[light.Y]++
		}
		for _, val := range ys {
			if val >= 5 {
				ysinarow++
			}
		}
		if ysinarow > 2 && p.LargestY-p.SmallestY > 5 && p.LargestY-p.SmallestY < 200 {
			width := 500
			height := 500
			upLeft := image.Point{0, 0}
			lowRight := image.Point{width, height}
			img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
			cyan := color.RGBA{100, 200, 200, 0xff}
			for _, light := range p.Lights {
				img.Set(light.X, light.Y, cyan)
			}
			f, _ := os.Create(fmt.Sprintf("image%d.png", i))
			png.Encode(f, img)

		}
		moveLights(&p)
	}
}

func moveLights(p *puzzle) {
	p.LargestX = 0
	p.SmallestX = 0
	p.SmallestY = 0
	p.LargestY = 0
	for _, light := range p.Lights {
		light.X += light.XVel
		light.Y += light.YVel
		if light.X > p.LargestX {
			p.LargestX = light.X
		}
		if light.Y > p.LargestY {
			p.LargestY = light.Y
		}
		if light.X < p.SmallestX {
			p.SmallestX = light.X
		}
		if light.Y < p.SmallestY {
			p.SmallestY = light.Y
		}
	}
}

func parseInput(data []string) puzzle {
	p := puzzle{
		LargestX: 0,
		LargestY: 0,
	}
	for _, row := range data {
		s := strings.Replace(row, " ", "", -1)
		s = strings.Replace(s, "<", "", -1)
		s = strings.Replace(s, ">", "", -1)
		s = strings.Replace(s, "position=", "", -1)
		s = strings.Replace(s, "velocity=", ",", -1)
		vals := strings.Split(s, ",")
		l := &light{
			X:    convertToInt(vals[0]),
			Y:    convertToInt(vals[1]),
			XVel: convertToInt(vals[2]),
			YVel: convertToInt(vals[3]),
		}
		if l.X > p.LargestX {
			p.LargestX = l.X
		}
		if l.Y > p.LargestY {
			p.LargestY = l.Y
		}
		if l.X < p.SmallestX {
			p.SmallestX = l.X
		}
		if l.Y < p.SmallestY {
			p.SmallestY = l.Y
		}
		p.Lights = append(p.Lights, l)
	}
	return p
}

func convertToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
