package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	world := World{
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
			NewCell(false),
			NewCell(false),

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
			NewCell(false),
			NewCell(false),
			NewCell(false),
		},
	}

	for true {
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

		world.resetScore().calcScore()
		fmt.Println(world)
		fmt.Print("enter: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		world.evalScore()
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
	score  int
}

func NewCell(isLive bool) Cell {
	return Cell{isLive, 0}
}

func (w World) calcScore() World {
	l := len(w.cells)

	for i, _ := range w.cells {
		uCellIdx := calcIndex(l, i, -w.col)
		if w.cells[uCellIdx].IsLive {
			w.cells[i].score += 1
		}
		ruCellIdx := calcIndex(l, i, -w.col+1)
		if w.cells[ruCellIdx].IsLive {
			w.cells[i].score += 1
		}
		rCellIdx := calcIndex(l, i, 1)
		if w.cells[rCellIdx].IsLive {
			w.cells[i].score += 1
		}
		rdCellIdx := calcIndex(l, i, w.col+1)
		if w.cells[rdCellIdx].IsLive {
			w.cells[i].score += 1
		}
		dCellIdx := calcIndex(l, i, w.col)
		if w.cells[dCellIdx].IsLive {
			w.cells[i].score += 1
		}
		ldCellIdx := calcIndex(l, i, w.col-1)
		if w.cells[ldCellIdx].IsLive {
			w.cells[i].score += 1
		}
		lCellIdx := calcIndex(l, i, -1)
		if w.cells[lCellIdx].IsLive {
			w.cells[i].score += 1
		}
		luCellIdx := calcIndex(l, i, -w.col-1)
		if w.cells[luCellIdx].IsLive {
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

func (w World) resetScore() World {
	for i, _ := range w.cells {
		w.cells[i].score = 0
	}

	return w
}

// 引数で渡す

// 方向関係なく、９つのマスの合計を数える
// 引数指定にする
