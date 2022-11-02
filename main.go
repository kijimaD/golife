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
		row: 3,
		col: 3,
		cells: []Cell{NewCell(false),
			NewCell(true),
			NewCell(false),

			NewCell(true),
			NewCell(true),
			NewCell(true),

			NewCell(false),
			NewCell(true),
			NewCell(false),
		},
	}

	for true {
		world.resetScore().calcScore()
		fmt.Println(world)
		fmt.Print("enter: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		world.evalScore()

		for i, c := range world.cells {
			if c.IsLive {
				fmt.Print("●")
			} else {
				fmt.Print("○")
			}

			if (i+1)%int(world.col) == 0 {
				fmt.Println("")
			}
		}
	}
}

// 全体
type World struct {
	row   int64
	col   int64
	cells []Cell
}

type Cell struct {
	IsLive bool
	score  int
}

func NewCell(isLive bool) Cell {
	return Cell{isLive, 0}
}

func (w World) calcScore() World {
	for i, _ := range w.cells {
		if w.CheckUpPos(int64(i + 1)) {
			w.cells[i].score += 1
		}
		if w.CheckRightPos(int64(i + 1)) {
			w.cells[i].score += 1
		}
		if w.CheckDownPos(int64(i + 1)) {
			w.cells[i].score += 1
		}
		if w.CheckLeftPos(int64(i + 1)) {
			w.cells[i].score += 1
		}
	}

	return w
}

func (w World) evalScore() World {
	for i, c := range w.cells {
		if c.score == 3 {
			w.cells[i].IsLive = true
		} else if c.score == 2 {
			w.cells[i].IsLive = true
		} else if c.score <= 1 {
			w.cells[i].IsLive = false
		} else if c.score >= 4 {
			w.cells[i].IsLive = false
		}
	}

	return w
}

// リセットできない
func (w World) resetScore() World {
	for i, _ := range w.cells {
		w.cells[i].score = 0
	}

	return w
}

func (w World) CheckLeftPos(i int64) bool {
	c := w.IntToCord(i)

	var tx int64
	if c.x == 1 {
		// 右端へ
		tx = w.col
	} else {
		tx = c.x - 1
	}

	newC := Cord{tx, c.y}

	// 1始まり
	return w.cells[newC.CordToInt(w.row, w.col)-1].IsLive
}

func (w World) CheckRightPos(i int64) bool {
	c := w.IntToCord(i)

	var tx int64
	if c.x == w.col {
		// 左端へ
		tx = 1
	} else {
		tx = c.x + 1
	}

	newC := Cord{tx, c.y}

	return w.cells[newC.CordToInt(w.row, w.col)-1].IsLive
}

func (w World) CheckUpPos(i int64) bool {
	c := w.IntToCord(i)

	var ty int64
	if c.y == 1 {
		// 下端へ
		ty = w.row
	} else {
		ty = c.y - 1
	}

	newC := Cord{c.x, ty}

	return w.cells[newC.CordToInt(w.row, w.col)-1].IsLive
}

func (w World) CheckDownPos(i int64) bool {
	c := w.IntToCord(i)

	var ty int64
	if c.y == w.row {
		// 上端へ
		ty = 1
	} else {
		ty = c.y + 1
	}

	newC := Cord{c.x, ty}

	return w.cells[newC.CordToInt(w.row, w.col)-1].IsLive
}

func (w World) IntToCord(idx int64) Cord {
	var cx int64
	var cy int64

	if int(idx) <= int(w.col) {
		cx = idx
	} else if idx%int64(w.col) == 0 {
		cx = w.col
	} else {
		cx = idx % w.col
	}

	if int(idx) <= int(w.row) {
		cy = 1
	} else if idx%w.row == 0 {
		cy = idx / w.row
	} else {
		cy = idx/w.row + 1
	}

	return Cord{x: cx, y: cy}
}

type Cord struct {
	x int64
	y int64
}

func (c *Cord) CordToInt(col int64, row int64) int64 {
	return (c.x-1)%col + 1 + ((c.y - 1) * row)
}
