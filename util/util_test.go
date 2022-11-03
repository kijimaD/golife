package util

import (
	"testing"
)

func TestCalcIndex(t *testing.T) {
	// 0 1 2
	// 3 4 5
	// 6 7 8

	// 4からの相対位置指定で、8方向分のインデックスを取得するテスト
	len := 9
	cur := 4
	col := 3

	tests := []struct {
		iw       int
		ix       int
		ip       int
		expected int
	}{
		{len, cur, 0, 4},        // 元の位置
		{len, cur, -col, 1},     // 上
		{len, cur, -col + 1, 2}, // 右上
		{len, cur, 1, 5},        // 右
		{len, cur, col + 1, 8},  // 右下
		{len, cur, col, 7},      // 下
		{len, cur, col - 1, 6},  // 左下
		{len, cur, -1, 3},       // 左
		{len, cur, -col - 1, 0}, // 左上
	}

	for i, tt := range tests {
		if calcIndex(tt.iw, tt.ix, tt.ip) != tt.expected {
			t.Errorf("idx: %d is not match", i)
		}
	}
}

func TestCalcIndex2(t *testing.T) {
	// 0  1  2  3  4
	// 5  6  7  8  9
	// 10 11 12 13 14
	// 15 16 17 18 19
	// 20 21 22 23 24

	// 4からの相対位置指定で、8方向分のインデックスを取得するテスト
	len := 25
	cur := 12
	col := 5

	tests := []struct {
		iw       int
		ix       int
		ip       int
		expected int
	}{
		{len, cur, 0, 12},       // 元の位置
		{len, cur, -col, 7},     // 上
		{len, cur, -col + 1, 8}, // 右上
		{len, cur, 1, 13},       // 右
		{len, cur, col + 1, 18}, // 右下
		{len, cur, col, 17},     // 下
		{len, cur, col - 1, 16}, // 左下
		{len, cur, -1, 11},      // 左
		{len, cur, -col - 1, 6}, // 左上
	}

	for i, tt := range tests {
		if calcIndex(tt.iw, tt.ix, tt.ip) != tt.expected {
			t.Errorf("idx: %d is not match", i)
		}
	}
}
