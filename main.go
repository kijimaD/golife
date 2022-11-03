package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	world := &World{
		row: 5,
		col: 5,
		cells: []Cell{
			NewCell(false),
			NewCell(false),
			NewCell(false),
			NewCell(false),
			NewCell(false),

			NewCell(false),
			NewCell(true),
			NewCell(true),
			NewCell(true),
			NewCell(false),

			NewCell(false),
			NewCell(false),
			NewCell(true),
			NewCell(true),
			NewCell(false),

			NewCell(false),
			NewCell(false),
			NewCell(false),
			NewCell(false),
			NewCell(false),

			NewCell(false),
			NewCell(false),
			NewCell(false),
			NewCell(false),
			NewCell(false),
		},
	}

	world.draw()
	gen := 0

	for true {
		fmt.Printf("%d gen: enter\n", gen)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		world.resetScore().calcScore().evalScore().draw()
		gen += 1
	}
}

// 全体
type World struct {
	row   int
	col   int
	cells []Cell
}

type Cell struct {
	IsLive bool
	Score  int
}

func NewCell(isLive bool) Cell {
	return Cell{isLive, 0}
}

func (w *World) calcScore() *World {
	l := len(w.cells)

	for i, _ := range w.cells {
		uCellIdx := calcIndex(l, i, -w.col)
		if w.cells[uCellIdx].IsLive {
			w.cells[i].Score += 1
		}
		ruCellIdx := calcIndex(l, i, -w.col+1)
		if w.cells[ruCellIdx].IsLive {
			w.cells[i].Score += 1
		}
		rCellIdx := calcIndex(l, i, 1)
		if w.cells[rCellIdx].IsLive {
			w.cells[i].Score += 1
		}
		rdCellIdx := calcIndex(l, i, w.col+1)
		if w.cells[rdCellIdx].IsLive {
			w.cells[i].Score += 1
		}
		dCellIdx := calcIndex(l, i, w.col)
		if w.cells[dCellIdx].IsLive {
			w.cells[i].Score += 1
		}
		ldCellIdx := calcIndex(l, i, w.col-1)
		if w.cells[ldCellIdx].IsLive {
			w.cells[i].Score += 1
		}
		lCellIdx := calcIndex(l, i, -1)
		if w.cells[lCellIdx].IsLive {
			w.cells[i].Score += 1
		}
		luCellIdx := calcIndex(l, i, -w.col-1)
		if w.cells[luCellIdx].IsLive {
			w.cells[i].Score += 1
		}
	}

	return w
}

func (w *World) evalScore() *World {
	for i, c := range w.cells {
		if c.Score == 3 { // 誕生
			w.cells[i].IsLive = true
		} else if c.Score == 2 && c.IsLive == true { // 生存
			w.cells[i].IsLive = true
		} else if c.Score <= 1 { // 過疎
			w.cells[i].IsLive = false
		} else if c.Score >= 4 { // 過密
			w.cells[i].IsLive = false
		}
	}

	return w
}

func (w *World) resetScore() *World {
	for i, _ := range w.cells {
		w.cells[i].Score = 0
	}
	return w
}

func (w World) draw() {
	for i, c := range w.cells {
		if c.IsLive {
			fmt.Print("●")
		} else {
			fmt.Print("○")
		}

		// 改行
		if (i+1)%int(w.col) == 0 {
			fmt.Println("")
		}
	}
}
