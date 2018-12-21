package main

import (
	"fmt"
	"testing"
)

func TestNumbersToAdd(t *testing.T) {
	tt := []struct {
		Elf1           int
		Elf2           int
		ExpectedResult []int
	}{
		{3, 7, []int{1, 0}},
		{6, 1, []int{1, 0}},
	}
	for _, tc := range tt {
		sum := numbersToAdd(tc.Elf1, tc.Elf2)
		fmt.Println(sum)
	}
}
