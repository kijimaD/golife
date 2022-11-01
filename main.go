package main

import (
	"fmt"
)

func main() {
	// ○●○
	// ●●●
	// ○●○
	world := World{
		Cell{1, false},
		Cell{2, true},
		Cell{3, false},

		Cell{4, true},
		Cell{5, true},
		Cell{6, true},

		Cell{7, false},
		Cell{8, true},
		Cell{9, false},
	}

	fmt.Println(world)
	fmt.Println(world.UpPos(5))

}

// 各セルをbooleanのスライスで表す
// [ true, true, false, true ]
// ●●
// ○●

const ROW = 3
const COL = 3

// 全体
type World []Cell

func (w World) LeftPos(i int64) Cell {
	// 通常
	// 3 -> 2
	// ループ
	// 4 -> 6

	fmt.Println(i, COL)
	if isLeftEdge(i) {
		fmt.Println("loop")
		return w[COL]
	} else {
		fmt.Println("hutu")
		return w[i-1]
	}
}

func isLeftEdge(i int64) bool {
	// 1, 4, 7...
	return i%COL == 1
}

func (w World) RightPos(i int64) Cell {
	// 最右のとき左に

	return w[i+1]
}

func (w World) UpPos(i int64) Cell {
	// 4 -> 1
	// 1 -> 7
	// 2 -> 5

	if int(i) <= ROW { // 上いっぱいのとき
		// 最後の列に移動
		upIdx := len(w) / ROW
		fmt.Println(upIdx)
		return w[upIdx]
	} else {
		return w[i+ROW]
	}
}

func (w World) DownPos(i int64) Cell {
	// 1 -> 4
	// 8 -> 2
	downIdx := int(i) + (len(w) / ROW) + 1

	if downIdx > len(w) { // 下いっぱいのとき
		// 最初の列に移動
		return Cell{int64(downIdx), false}
	} else {
		return w[downIdx]
	}
}

func (w World) IntToCord(i int64) []int64 {
	// 1 -> (1,1)
	// 2 -> (2,1)
	// 3 -> (3,1)
	// 4 -> (1,2)
	// 5 -> (2,2)
	// 6 -> (3,2)

	var cx int64
	var cy int64

	if int(i) <= COL {
		cx = i
	} else if i%COL == 0 {
		cx = COL
	} else {
		cx = i % COL
	}

	if int(i) <= ROW {
		cy = 1
	} else if i%ROW == 0 {
		cy = i / ROW
	} else {
		cy = i/ROW + 1
	}

	return []int64{cx, cy}
}

// セル１つ１つ
type Cell struct {
	index  int64
	IsLive bool
}
