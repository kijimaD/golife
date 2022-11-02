package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// ○●○
	// ●●●
	// ○●○
	world := World{
		NewCell(false),
		NewCell(true),
		NewCell(false),

		NewCell(true),
		NewCell(true),
		NewCell(true),

		NewCell(false),
		NewCell(true),
		NewCell(false),
	}

	for true {
		fmt.Println(world)
		fmt.Print("enter: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		world.calcScore()
	}
}

const ROW = 3
const COL = 3

// 全体
type World []Cell

type Cell struct {
	IsLive bool
	score  int
}

func NewCell(isLive bool) Cell {
	return Cell{isLive, 0}
}

func (w World) calcScore() World {
	for i, _ := range w {
		if w.CheckUpPos(int64(i + 1)) {
			w[i].score += 1
		}
		if w.CheckRightPos(int64(i + 1)) {
			w[i].score += 1
		}
		if w.CheckDownPos(int64(i + 1)) {
			w[i].score += 1
		}
		if w.CheckLeftPos(int64(i + 1)) {
			w[i].score += 1
		}
	}

	return w
}

func (w World) CheckLeftPos(i int64) bool {
	c := IntToCord(i)

	var tx int64
	if c.x == 1 {
		// 右端へ
		tx = COL
	} else {
		tx = c.x - 1
	}

	newC := Cord{tx, c.y}

	// 1始まり
	return w[newC.CordToInt()-1].IsLive
}

func (w World) CheckRightPos(i int64) bool {
	c := IntToCord(i)

	var tx int64
	if c.x == COL {
		// 左端へ
		tx = 1
	} else {
		tx = c.x + 1
	}

	newC := Cord{tx, c.y}

	return w[newC.CordToInt()-1].IsLive
}

func (w World) CheckUpPos(i int64) bool {
	c := IntToCord(i)

	var ty int64
	if c.y == 1 {
		// 下端へ
		ty = ROW
	} else {
		ty = c.y - 1
	}

	newC := Cord{c.x, ty}

	return w[newC.CordToInt()-1].IsLive
}

func (w World) CheckDownPos(i int64) bool {
	c := IntToCord(i)

	var ty int64
	if c.y == ROW {
		// 上端へ
		ty = 1
	} else {
		ty = c.y + 1
	}

	newC := Cord{c.x, ty}

	return w[newC.CordToInt()-1].IsLive
}

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

type Cord struct {
	x int64
	y int64
}

func (c *Cord) CordToInt() int64 {
	return (c.x-1)%ROW + 1 + ((c.y - 1) * COL)
}
