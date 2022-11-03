package main

func calcIndex(w int, x int, r int) int {
	// x: 基準になるインデックス
	// r: 相対的な位置
	nx := (x + r) % w
	if nx < 0 {
		nx += w
	}

	return nx
}
