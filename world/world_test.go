package world

import (
	"testing"
)

func worldSeed() *World {
	// ○○○○○
	// ○○●○○
	// ○●○●○
	// ○○●○○
	// ○○○○○
	return &World{
		cells: []Cell{
			NewCell(false),
			NewCell(false),
			NewCell(false),
			NewCell(false),
			NewCell(false),

			NewCell(false),
			NewCell(false),
			NewCell(true),
			NewCell(false),
			NewCell(false),

			NewCell(false),
			NewCell(true),
			NewCell(false),
			NewCell(true),
			NewCell(false),

			NewCell(false),
			NewCell(false),
			NewCell(true),
			NewCell(false),
			NewCell(false),

			NewCell(false),
			NewCell(false),
			NewCell(false),
			NewCell(false),
			NewCell(false),
		},
		row: 5,
		col: 5,
	}
}

func TestCalcScore(t *testing.T) {
	w := worldSeed()

	tests := []struct {
		expected int
	}{
		{0},
		{1},
		{1},
		{1},
		{0},

		{1},
		{2},
		{2},
		{2},
		{1},

		{1},
		{2},
		{4},
		{2},
		{1},

		{1},
		{2},
		{2},
		{2},
		{1},

		{0},
		{1},
		{1},
		{1},
		{0},
	}

	w.calcScore()

	for i, tt := range tests {
		if w.cells[i].Score != tt.expected {
			t.Errorf("input: %v is not match. returned: %d, expected: %d", i, w.cells[i].Score, tt.expected)
		}
	}
}

func TestResetScore(t *testing.T) {
	w := worldSeed()

	tests := []struct {
		expected int
	}{
		{0},
		{0},
		{0},

		{0},
		{0},
		{0},

		{0},
		{0},
		{0},
	}

	w.resetScore()

	for i, tt := range tests {
		if w.cells[i].Score != tt.expected {
			t.Errorf("input: %v is not match", i)
		}
	}
}
