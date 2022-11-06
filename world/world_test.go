package world

import (
	"golife/config"
	"testing"
)

func worldSeed() *World {
	// ○○○○○
	// ○○●○○
	// ○●○●○
	// ○○●○○
	// ○○○○○
	return &World{
		Cells: []Cell{
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
		configs: config.Configs{
			Row: 5,
			Col: 5,
		},
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
		if w.Cells[i].Score != tt.expected {
			t.Errorf("input: %v is not match. returned: %d, expected: %d", i, w.Cells[i].Score, tt.expected)
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
		if w.Cells[i].Score != tt.expected {
			t.Errorf("input: %v is not match", i)
		}
	}
}
