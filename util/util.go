package util

var (
	// structを定数にできないのでvarにしてる
	Current   = coord{x: 0, y: 0}
	Up        = coord{x: 0, y: -1}
	RightUp   = coord{x: 1, y: -1}
	Right     = coord{x: 1, y: 0}
	RightDown = coord{x: 1, y: 1}
	Down      = coord{x: 0, y: 1}
	LeftDown  = coord{x: -1, y: 1}
	Left      = coord{x: -1, y: 0}
	LeftUp    = coord{x: -1, y: -1}
)

type coord struct {
	x int
	y int
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

// 2次元配列で相対移動したインデックスを求める
// w: 2次元配列の横と縦の長さ。今のところ正方形しか対応してない
// i: 基準になるインデックス
// r: 相対的な位置
func PlaneIndex(w int, i int, r coord) int {
	x := CalcIndex(w, i%w, r.x)
	y := CalcIndex(w, i/w, r.y)

	return x + y*w
}

// ほんとはserverに置きたいけど、serverと循環参照が起きてしまうのでここに置くことにした...
type CreateWorldParams struct {
	InitialWorld string `form:"InitialWorld"`
	GenCap       int    `form:"GenCap"`
	Debug        bool   `form:"Debug"`
}
