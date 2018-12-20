package main

import (
	"testing"
)

func TestNextState(t *testing.T) {
	notes := []note{
		note{
			Pattern: "...##",
			NextGen: "#",
		},
		note{
			Pattern: "..#..",
			NextGen: "#",
		},
		note{
			Pattern: ".#...",
			NextGen: "#",
		},
		note{
			Pattern: ".#.#.",
			NextGen: "#",
		},
		note{
			Pattern: ".#.##",
			NextGen: "#",
		},
		note{
			Pattern: ".##..",
			NextGen: "#",
		},
		note{
			Pattern: ".####",
			NextGen: "#",
		},
		note{
			Pattern: "#.#.#",
			NextGen: "#",
		},
		note{
			Pattern: "#.###",
			NextGen: "#",
		},
		note{
			Pattern: "##.#.",
			NextGen: "#",
		},
		note{
			Pattern: "##.##",
			NextGen: "#",
		},
		note{
			Pattern: "###..",
			NextGen: "#",
		},
		note{
			Pattern: "###.#",
			NextGen: "#",
		},
		note{
			Pattern: "####.",
			NextGen: "#",
		},
	}
	tt := []struct {
		currentState  string
		expectedState string
	}{

		{"#..#.#..##......###...###", "...#...#....#.....#..#..#..#..."},
		{"...#...#....#.....#..#..#..#...........", "...##..##...##....#..#..#..##.........."},
	}
	for index, tc := range tt {
		s, _ := genNextState(tc.currentState, notes)
		if s != tc.expectedState {
			t.Errorf("\n%d: Tested\t'%s'\nExpected\t'%s'\nGot\t\t'%s'", index+1, tc.currentState, tc.expectedState, s)
		}
	}
}

func TestCalcScore(t *testing.T) {
	tt := []struct {
		currentState  string
		indexShift    int
		expectedScore int
	}{
		{".#....##....#####...#######....#.#..##.", -3, 325},
		{"...#....##....#####...#######....#.#..##..", -5, 325},
	}
	for index, tc := range tt {
		score := calculatePotScore(tc.currentState, tc.indexShift)
		if score != tc.expectedScore {
			t.Errorf("%d: Expected: %d, Got: %d", index, tc.expectedScore, score)
		}
	}
}
