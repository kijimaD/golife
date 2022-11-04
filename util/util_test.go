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
		rx       int
		ry       int
		expected int
	}{
		{col, cur, 0, 0, 4},   // 元の位置
		{col, cur, 0, -1, 1},  // 上
		{col, cur, 1, -1, 2},  // 右上
		{col, cur, 1, 0, 5},   // 右
		{col, cur, 1, 1, 8},   // 右下
		{col, cur, 0, 1, 7},   // 下
		{col, cur, -1, 1, 6},  // 左下
		{col, cur, -1, 0, 3},  // 左
		{col, cur, -1, -1, 0}, // 左上
	}

	for i, tt := range tests {
		if PlaneIndex(tt.iw, tt.base, tt.rx, tt.ry) != tt.expected {
			t.Errorf("idx: %d is not match", i)
		}
	}
}

func TestPlaneIndexEdge(t *testing.T) {
	// 0 1 2
	// 3 4 5
	// 6 7 8

	// 0からの相対位置指定で、8方向分のインデックスを取得するテスト。トーラス処理を確かめる
	cur := 0
	col := 3

	tests := []struct {
		iw       int
		base     int
		rx       int
		ry       int
		expected int
	}{
		{col, cur, 0, 0, 0},   // 元の位置
		{col, cur, 0, -1, 6},  // 上
		{col, cur, 1, -1, 7},  // 右上
		{col, cur, 1, 0, 1},   // 右
		{col, cur, 1, 1, 4},   // 右下
		{col, cur, 0, 1, 3},   // 下
		{col, cur, -1, 1, 5},  // 左下
		{col, cur, -1, 0, 2},  // 左
		{col, cur, -1, -1, 8}, // 左上
	}

	for i, tt := range tests {
		if PlaneIndex(tt.iw, tt.base, tt.rx, tt.ry) != tt.expected {
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
