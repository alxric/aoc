package main

import "testing"

func TestPowerLevel(t *testing.T) {
	if val := calculatePowerLevel(1, 1, 6548); val != -4 {
		t.Errorf("Expected -4, got %d", val)
	}
	if val := calculatePowerLevel(1, 2, 6548); val != -3 {
		t.Errorf("Expected -3, got %d", val)
	}
	if val := calculatePowerLevel(2, 1, 6548); val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}
	if val := calculatePowerLevel(2, 2, 6548); val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}
}

func TestSquare(t *testing.T) {
	p := generateGrid(300, 18)
	if val := buildSquare(&p, 89, 268, 16); val != 113 {
		t.Errorf("Expected 113, got %d", val)
	}
	p = generateGrid(300, 42)
	if val := buildSquare(&p, 231, 250, 12); val != 119 {
		t.Errorf("Expected 119, got %d", val)
	}
}
