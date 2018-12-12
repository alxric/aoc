package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	s := parseInput()
	two, three := countLetters(s)
	var wg sync.WaitGroup
	twoSum, threeSum := 0, 0
	wg.Add(2)
	go func() {
		defer wg.Done()
		for num := range two {
			twoSum += num
		}
	}()
	go func() {
		defer wg.Done()
		for num := range three {
			threeSum += num
		}
	}()
	wg.Wait()
	fmt.Println(twoSum * threeSum)
}

func countLetters(s <-chan string) (<-chan int, <-chan int) {
	two, three := make(chan int), make(chan int)
	go func() {
		for row := range s {
			twoSent, threeSent := false, false
			m := make(map[string]int)
			for _, char := range row {
				if _, ok := m[string(char)]; !ok {
					m[string(char)] = 0
				}
				m[string(char)]++
			}

			for _, count := range m {
				switch count {
				case 2:
					if !twoSent {
						two <- 1
						three <- 0
					}
					twoSent = true
				case 3:
					if !threeSent {
						two <- 0
						three <- 1
					}
					threeSent = true
				}
			}

		}
		close(two)
		close(three)
	}()
	return two, three
}

func parseInput() <-chan string {
	s := make(chan string)
	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		for scanner.Scan() {
			s <- scanner.Text()
		}
		close(s)

	}()
	return s
}
