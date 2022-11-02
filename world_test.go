package main

import (
	"testing"
)

func worldSeed() World {
	// ○●○
	// ●●●
	// ○●○
	return World{
		Cell{false},
		Cell{true},
		Cell{false},

		Cell{true},
		Cell{true},
		Cell{true},

		Cell{false},
		Cell{true},
		Cell{false},
	}
}

func TestCheckLeftPos(t *testing.T) {
	w := worldSeed()

	tests := []struct {
		input    int64
		expected bool
	}{
		{1, false},
		{2, false},
		{3, true},
		{4, true},
		{5, true},
		{6, true},
	}

	for _, tt := range tests {
		if w.CheckLeftPos(tt.input) != tt.expected {
			t.Errorf("input: %v is not match", tt.input)
		}
	}
}

func TestCheckUpPos(t *testing.T) {
	// 通常
	// if worldSeed().CheckUpPos(4) != worldSeed()[1] {
	// 	t.Errorf("is not match")
	// }

	// // 乗り越え
	// if worldSeed().CheckUpPos(1) != worldSeed()[7] {
	// 	t.Errorf("is not match")
	// }
}

func TestIntToCord(t *testing.T) {
	tests := []struct {
		input    int64
		expected Cord
	}{
		{1, Cord{1, 1}},
		{2, Cord{2, 1}},
		{3, Cord{3, 1}},
		{4, Cord{1, 2}},
		{5, Cord{2, 2}},
		{6, Cord{3, 2}},
		{7, Cord{1, 3}},
		{8, Cord{2, 3}},
		{9, Cord{3, 3}},
	}

	for _, tt := range tests {
		if c := IntToCord(tt.input); c != tt.expected {
			t.Errorf("%v: %v is not match, expected: %v", tt.input, c, tt.expected)
		}
	}
}

func TestCordToInt(t *testing.T) {
	tests := []struct {
		input    Cord
		expected int64
	}{
		{Cord{1, 1}, 1},
		{Cord{2, 1}, 2},
		{Cord{3, 1}, 3},
		{Cord{1, 2}, 4},
		{Cord{2, 2}, 5},
		{Cord{3, 2}, 6},
		{Cord{1, 3}, 7},
		{Cord{2, 3}, 8},
		{Cord{3, 3}, 9},
	}

	for _, tt := range tests {
		if i := tt.input.CordToInt(); i != tt.expected {
			t.Errorf("%v: %v is not match, expected: %v", tt.input, i, tt.expected)
		}
	}
}
