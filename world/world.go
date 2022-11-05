package world

import (
	"bufio"
	"fmt"
	"golife/util"
	"io"
	"io/ioutil"
	"os"
)

const INPUT_FILE = "world.txt"

// 全体
type World struct {
	row   int
	col   int
	cells []Cell
}

// TODO: 今はファイルだけ。readerで読み込めるようにする
func LoadWorld() *World {
	var row int
	var col int

	f, _ := os.Open(INPUT_FILE)
	bu := bufio.NewReaderSize(f, 1024)
	for {
		line, _, err := bu.ReadLine()
		if err == io.EOF {
			break
		}
		// 最終行の文字長をcolとする
		col = len([]rune(string(line)))
		// 行数をrowとする
		row += 1
	}
	f.Close()

	w := &World{
		row: row,
		col: col,
	}

	data, _ := ioutil.ReadFile(INPUT_FILE)
	for _, c := range string(data) {
		switch string(c) {
		case LIVEC:
			w.cells = append(w.cells, NewCell(true))
		case DEATHC:
			w.cells = append(w.cells, NewCell(false))
		case "\n":
		default:
			fmt.Printf("`%s`は不正な文字です\n", string(c))
		}
	}

	return w
}

func (w *World) CalcScore() *World {
	for i, _ := range w.cells {
		if w.cells[util.PlaneIndex(w.col, i, util.Up)].IsLive {
			w.cells[i].Score += 1
		}
		if w.cells[util.PlaneIndex(w.col, i, util.RightUp)].IsLive {
			w.cells[i].Score += 1
		}
		if w.cells[util.PlaneIndex(w.col, i, util.Right)].IsLive {
			w.cells[i].Score += 1
		}
		if w.cells[util.PlaneIndex(w.col, i, util.RightDown)].IsLive {
			w.cells[i].Score += 1
		}
		if w.cells[util.PlaneIndex(w.col, i, util.Down)].IsLive {
			w.cells[i].Score += 1
		}
		if w.cells[util.PlaneIndex(w.col, i, util.LeftDown)].IsLive {
			w.cells[i].Score += 1
		}
		if w.cells[util.PlaneIndex(w.col, i, util.Left)].IsLive {
			w.cells[i].Score += 1
		}
		if w.cells[util.PlaneIndex(w.col, i, util.LeftUp)].IsLive {
			w.cells[i].Score += 1
		}
	}

	return w
}

func (w *World) EvalScore() *World {
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

func (w *World) ResetScore() *World {
	for i, _ := range w.cells {
		w.cells[i].Score = 0
	}
	return w
}

func (w World) Draw() {
	for i, c := range w.cells {
		fmt.Print(c.String())
		// 改行
		if (i+1)%int(w.col) == 0 {
			fmt.Println("")
		}
	}
}
