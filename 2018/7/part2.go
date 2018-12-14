package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type step struct {
	Name         string
	Dependencies map[string]bool
}

type worker struct {
	ID              int
	WorkingWith     string
	Working         bool
	ExpectedSeconds int
	WorkedSeconds   int
}

func main() {
	start := make(chan bool)
	startPrint := make(chan bool)
	done := make(chan bool)
	completed := make(chan string)
	workChan := make(chan string)
	workers := buildWorkers()
	m := parseInput()
	sendNextItem(m, start, workChan, completed, done)
	doWork(m, workers, workChan, startPrint)
	printStatus(workers, completed, startPrint)
	time.Sleep(time.Second * 1)
	start <- true
	<-done
}

func printStatus(workers []*worker, completed chan string, startPrint chan bool) {
	var finalString string
	go func() {
		for {
			select {
			case <-startPrint:
				runTimer := 1
				fmt.Println("Second\tWorker 1\tWorker 2\tWorker 3\tWorker 4\tWorker 5\tDone")
				for {
					fmt.Println(
						fmt.Sprintf("%d\t      %s\t\t      %s\t\t      %s\t\t      %s\t\t      %s\t%s",
							runTimer, workers[0].WorkingWith, workers[1].WorkingWith, workers[2].WorkingWith,
							workers[3].WorkingWith, workers[4].WorkingWith, finalString),
					)

					for _, worker := range workers {
						if worker.Working {
							worker.WorkedSeconds++
							if worker.WorkedSeconds == worker.ExpectedSeconds {
								worker.Working = false
								completed <- worker.WorkingWith
								finalString += worker.WorkingWith
								worker.WorkingWith = ""
							}
						}
					}
					runTimer++
					time.Sleep(time.Millisecond * 1)
				}
			}
		}
	}()
}

func doWork(m map[string]*step, workers []*worker, workChan chan string, startPrint chan bool) {
	started := false
	go func() {
		for {
			select {
			case wo := <-workChan:
				for _, worker := range workers {
					if !worker.Working {
						worker.Working = true
						worker.WorkingWith = wo
						worker.WorkedSeconds = 0
						worker.ExpectedSeconds = 60 + int([]rune(wo)[0]) - 64
						if !started {
							started = true
							startPrint <- true
						}
						break
					}
				}
			}
		}
	}()
}

func sendNextItem(m map[string]*step, start chan bool, workChan chan string, completed chan string, done chan bool) {
	go func() {
		for {
			select {
			case <-start:
				wo := workOrder(m)
				for _, letter := range wo {
					delete(m, letter)
					workChan <- letter
				}
			case l := <-completed:
				for _, s := range m {
					delete(s.Dependencies, l)
				}
				wo := workOrder(m)
				if len(wo) == 0 && len(m) == 0 {
					done <- true
				}
				for _, letter := range wo {
					delete(m, letter)
					workChan <- letter
				}
			}
		}

	}()

}

func buildWorkers() []*worker {
	var workers []*worker
	for i := 1; i <= 5; i++ {
		w := &worker{
			ID: i,
		}
		workers = append(workers, w)
	}
	return workers
}

func workOrder(m map[string]*step) []string {
	var doNow []string
	for _, s := range m {
		if len(s.Dependencies) == 0 {
			doNow = append(doNow, s.Name)
		}
	}
	sort.Strings(doNow)
	return doNow
}

func parseInput() map[string]*step {
	m := make(map[string]*step)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		if _, ok := m[s[1]]; !ok {
			m[s[1]] = &step{
				Name:         s[1],
				Dependencies: map[string]bool{},
			}
		}
		if ss, ok := m[s[7]]; !ok {
			m[s[7]] = &step{
				Name: s[7],
				Dependencies: map[string]bool{
					s[1]: true,
				},
			}
		} else {
			ss.Dependencies[s[1]] = true
		}
	}
	return m
}
