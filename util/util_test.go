package util

import (
	"testing"
)

func TestPlaneIndex(t *testing.T) {
	// 0 1 2
	// 3 4 5
	// 6 7 8

	// 4からの相対位置指定で、8方向分のインデックスを取得するテスト
	cur := 4
	col := 3

	tests := []struct {
		iw       int
		base     int
		r        coord
		expected int
	}{
		{col, cur, Current, 4},
		{col, cur, Up, 1},
		{col, cur, RightUp, 2},
		{col, cur, Right, 5},
		{col, cur, RightDown, 8},
		{col, cur, Down, 7},
		{col, cur, LeftDown, 6},
		{col, cur, Left, 3},
		{col, cur, LeftUp, 0},
	}

	for i, tt := range tests {
		if PlaneIndex(tt.iw, tt.base, tt.r) != tt.expected {
			t.Errorf("idx: %d is not match", i)
		}
	}
}

func TestPlaneIndexEdge0(t *testing.T) {
	// 0 1 2
	// 3 4 5
	// 6 7 8

	// 0からの相対位置指定で、8方向分のインデックスを取得するテスト。トーラス処理を確かめる
	cur := 0
	col := 3

	tests := []struct {
		iw       int
		base     int
		r        coord
		expected int
	}{
		{col, cur, Current, 0},
		{col, cur, Up, 6},
		{col, cur, RightUp, 7},
		{col, cur, Right, 1},
		{col, cur, RightDown, 4},
		{col, cur, Down, 3},
		{col, cur, LeftDown, 5},
		{col, cur, Left, 2},
		{col, cur, LeftUp, 8},
	}

	for i, tt := range tests {
		if PlaneIndex(tt.iw, tt.base, tt.r) != tt.expected {
			t.Errorf("idx: %d is not match", i)
		}
	}
}

func TestPlaneIndexEdge1(t *testing.T) {
	// 0 1 2
	// 3 4 5
	// 6 7 8

	// 5からの相対位置指定で、8方向分のインデックスを取得するテスト。トーラス処理を確かめる
	cur := 5
	col := 3

	tests := []struct {
		iw       int
		base     int
		r        coord
		expected int
	}{
		{col, cur, Current, 5},
		{col, cur, Up, 2},
		{col, cur, RightUp, 0},
		{col, cur, Right, 3},
		{col, cur, RightDown, 6},
		{col, cur, Down, 8},
		{col, cur, LeftDown, 7},
		{col, cur, Left, 4},
		{col, cur, LeftUp, 1},
	}

	for i, tt := range tests {
		if r := PlaneIndex(tt.iw, tt.base, tt.r); r != tt.expected {
			t.Errorf("idx: %d is not match. result: %d, expected: %d", i, r, tt.expected)
		}
	}
}

func TestPlaneIndexEdge2(t *testing.T) {
	// 0 1 2
	// 3 4 5
	// 6 7 8

	// 1からの相対位置指定で、8方向分のインデックスを取得するテスト。トーラス処理を確かめる
	cur := 1
	col := 3

	tests := []struct {
		iw       int
		base     int
		r        coord
		expected int
	}{
		{col, cur, Current, 1},
		{col, cur, Up, 7},
		{col, cur, RightUp, 8},
		{col, cur, Right, 2},
		{col, cur, RightDown, 5},
		{col, cur, Down, 4},
		{col, cur, LeftDown, 3},
		{col, cur, Left, 0},
		{col, cur, LeftUp, 6},
	}

	for i, tt := range tests {
		if r := PlaneIndex(tt.iw, tt.base, tt.r); r != tt.expected {
			t.Errorf("idx: %d is not match. result: %d, expected: %d", i, r, tt.expected)
		}
	}
}

func TestPlaneIndexBig(t *testing.T) {
	// 0  1  2  3  4
	// 5  6  7  8  9
	// 10 11 12 13 14
	// 15 16 17 18 19
	// 20 21 22 23 24

	// 24からの相対位置指定で、8方向分のインデックスを取得するテスト。
	cur := 24
	col := 5

	tests := []struct {
		iw       int
		base     int
		r        coord
		expected int
	}{
		{col, cur, Current, 24},
		{col, cur, Up, 19},
		{col, cur, RightUp, 15},
		{col, cur, Right, 20},
		{col, cur, RightDown, 0},
		{col, cur, Down, 4},
		{col, cur, LeftDown, 3},
		{col, cur, Left, 23},
		{col, cur, LeftUp, 18},
	}

	for i, tt := range tests {
		if PlaneIndex(tt.iw, tt.base, tt.r) != tt.expected {
			t.Errorf("idx: %d is not match", i)
		}
	}
}

func TestCalcIndex(t *testing.T) {
	// 0 1 2

	// 0からの相対位置指定で、2方向分のインデックスを取得するテスト
	cur := 0
	col := 3

	tests := []struct {
		iw       int
		ix       int
		ip       int
		expected int
	}{
		{col, cur, 0, 0},  // 元の位置
		{col, cur, 1, 1},  // 右
		{col, cur, -1, 2}, // 左
	}

	for i, tt := range tests {
		if r := CalcIndex(tt.iw, tt.ix, tt.ip); r != tt.expected {
			t.Errorf("idx: %d is not match. result: %d, expected: %d", i, r, tt.expected)
		}
	}
}
