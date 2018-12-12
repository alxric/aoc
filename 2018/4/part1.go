package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type guardInfo struct {
	ID                int
	TotalMinutesSlept float64
	StartsSleeping    time.Time
	MinuteCount       map[int]int
}

func main() {
	ts, m := sortInput()
	guardStats := genGuardStats(ts, m)
	guard := sleepiestGuard(guardStats)
	bestMinute := 0
	maxCount := 0
	for minute, count := range guard.MinuteCount {
		if count > maxCount {
			maxCount = count
			bestMinute = minute
		}
	}
	fmt.Println(guard.ID * bestMinute)
}

func sleepiestGuard(m map[int]*guardInfo) *guardInfo {
	maxSleep := float64(0)
	guard := 0
	for index, g := range m {
		if g.TotalMinutesSlept > maxSleep {
			guard = index
			maxSleep = g.TotalMinutesSlept
		}
	}
	return m[guard]
}

func genGuardStats(ts []string, m map[string]string) map[int]*guardInfo {
	guardInfos := make(map[int]*guardInfo)
	layout := "2006-01-02 15:04"
	currentGuard := 0
	for _, s := range ts {
		switch {
		case strings.Contains(m[s], "Guard"):
			words := strings.Split(m[s], " ")
			guardID := words[1][1:]
			iGuardID, err := strconv.Atoi(guardID)
			if err != nil {
				panic(err)
			}
			if _, ok := guardInfos[iGuardID]; !ok {
				g := &guardInfo{
					ID:          iGuardID,
					MinuteCount: make(map[int]int),
				}
				guardInfos[iGuardID] = g
			}
			currentGuard = iGuardID
		case strings.Contains(m[s], "falls asleep"):
			t, err := time.Parse(layout, s)
			if err != nil {
				panic(err)
			}
			guardInfos[currentGuard].StartsSleeping = t
		case strings.Contains(m[s], "wakes up"):
			t, err := time.Parse(layout, s)
			if err != nil {
				panic(err)
			}
			for i := guardInfos[currentGuard].StartsSleeping.Minute(); i < t.Minute(); i++ {
				guardInfos[currentGuard].MinuteCount[i]++
			}
			diff := t.Sub(guardInfos[currentGuard].StartsSleeping)
			guardInfos[currentGuard].TotalMinutesSlept += diff.Minutes()
		}
	}
	return guardInfos
}

func sortInput() ([]string, map[string]string) {
	scanner := bufio.NewScanner(os.Stdin)
	ts := []string{}
	m := make(map[string]string)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "]")
		t := s[0][1:]
		ts = append(ts, t)
		m[t] = strings.TrimSpace(s[1])
	}
	sort.Strings(ts)
	return ts, m
}
