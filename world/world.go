package world

import (
	"fmt"
	"golife/util"
)

// 全体
type World struct {
	Row   int
	Col   int
	Cells []Cell
}

type Cell struct {
	IsLive bool
	Score  int
}

func NewCell(isLive bool) Cell {
	return Cell{isLive, 0}
}

func (w *World) CalcScore() *World {
	l := len(w.Cells)

	for i, _ := range w.Cells {
		uCellIdx := util.CalcIndex(l, i, -w.Col)
		if w.Cells[uCellIdx].IsLive {
			w.Cells[i].Score += 1
		}
		ruCellIdx := util.CalcIndex(l, i, -w.Col+1)
		if w.Cells[ruCellIdx].IsLive {
			w.Cells[i].Score += 1
		}
		rCellIdx := util.CalcIndex(l, i, 1)
		if w.Cells[rCellIdx].IsLive {
			w.Cells[i].Score += 1
		}
		rdCellIdx := util.CalcIndex(l, i, w.Col+1)
		if w.Cells[rdCellIdx].IsLive {
			w.Cells[i].Score += 1
		}
		dCellIdx := util.CalcIndex(l, i, w.Col)
		if w.Cells[dCellIdx].IsLive {
			w.Cells[i].Score += 1
		}
		ldCellIdx := util.CalcIndex(l, i, w.Col-1)
		if w.Cells[ldCellIdx].IsLive {
			w.Cells[i].Score += 1
		}
		lCellIdx := util.CalcIndex(l, i, -1)
		if w.Cells[lCellIdx].IsLive {
			w.Cells[i].Score += 1
		}
		luCellIdx := util.CalcIndex(l, i, -w.Col-1)
		if w.Cells[luCellIdx].IsLive {
			w.Cells[i].Score += 1
		}
	}

	return w
}

func (w *World) EvalScore() *World {
	for i, c := range w.Cells {
		if c.Score == 3 { // 誕生
			w.Cells[i].IsLive = true
		} else if c.Score == 2 && c.IsLive == true { // 生存
			w.Cells[i].IsLive = true
		} else if c.Score <= 1 { // 過疎
			w.Cells[i].IsLive = false
		} else if c.Score >= 4 { // 過密
			w.Cells[i].IsLive = false
		}
	}

	return w
}

func (w *World) ResetScore() *World {
	for i, _ := range w.Cells {
		w.Cells[i].Score = 0
	}
	return w
}

func (w World) Draw() {
	for i, c := range w.Cells {
		if c.IsLive {
			fmt.Print("●")
		} else {
			fmt.Print("○")
		}

		// 改行
		if (i+1)%int(w.Col) == 0 {
			fmt.Println("")
		}
	}
}
