package main

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
	// 0  1  2  3  4
	// 5  6  7  8  9
	// 10 11 12 13 14
	// 15 16 17 18 19
	// 20 21 22 23 24

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

// ９つのマスを数を数える
// 方向は関係なく

// 畳み込み?

// 向こう側にいく処理の条件分岐をなくす
// -インデックスを使う? 長さ３の配列だと、-1 == 3 みたいな
