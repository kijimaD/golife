package util

// 2次元配列で相対移動したインデックスを求める
// w: 2次元配列の横と縦の長さ。今のところ正方形しか対応してない
// base: 基準になるインデックス
// rx: xの相対的な位置
// ry: yの相対的な位置
func PlaneIndex(w int, base int, rx int, ry int) int {
	x := CalcIndex(w, base, rx)
	y := CalcIndex(w, base, ry)

	return x + y*w
}

// 1次元配列で相対移動したインデックスを求める
// w: 実配列の長さ
// x: 基準になるインデックス
// r: 相対的な位置
func CalcIndex(w int, x int, r int) int {
	nx := (x + r) % w
	if nx < 0 {
		nx += w
	}

	return nx
}
