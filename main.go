package main

import (
	"fmt"
)

func main() {
	// ○●○
	// ●●●
	// ○●○
	world := World{
		Cell{false},
		Cell{true},
		Cell{false},

		Cell{true},
		Cell{true},
		Cell{true},

		Cell{false},
		Cell{true},
		Cell{false},
	}

	fmt.Println(world)
}

// 各セルをbooleanのスライスで表す
// [ true, true, false, true ]
// ●●
// ○●

const ROW = 3
const COL = 3

// 全体
type World []Cell

// func (w World) CheckLeftPos(i int64) Cell {

// }

// func (w World) CheckRightPos(i int64) Cell {

// }

// func (w World) CheckUpPos(i int64) Cell {

// }

// func (w World) CheckDownPos(i int64) Cell {
// }

func IntToCord(idx int64) Cord {
	var cx int64
	var cy int64

	if int(idx) <= COL {
		cx = idx
	} else if idx%COL == 0 {
		cx = COL
	} else {
		cx = idx % COL
	}

	if int(idx) <= ROW {
		cy = 1
	} else if idx%ROW == 0 {
		cy = idx / ROW
	} else {
		cy = idx/ROW + 1
	}

	return Cord{x: cx, y: cy}
}

// セル１つ１つ
type Cell struct {
	IsLive bool
}

type Cord struct {
	x int64
	y int64
}
