package world

import (
	"fmt"
	"golife/config"
	"golife/util"
	"io/ioutil"
)

const INPUT_FILE = "world.txt"

// 全体
type World struct {
	cells   []Cell
	Configs config.Configs
}

// TODO: 今はファイルだけ。readerで読み込めるようにする
func Load(c config.Configs) *World {
	w := &World{
		Configs: c,
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

func (w World) Next() World {
	return w.ResetScore().CalcScore().EvalScore()
}

func (w World) CalcScore() World {
	for i, _ := range w.cells {
		if w.cells[util.PlaneIndex(w.Configs.Col, i, util.Up)].IsLive {
			w.cells[i].Score += 1
		}
		if w.cells[util.PlaneIndex(w.Configs.Col, i, util.RightUp)].IsLive {
			w.cells[i].Score += 1
		}
		if w.cells[util.PlaneIndex(w.Configs.Col, i, util.Right)].IsLive {
			w.cells[i].Score += 1
		}
		if w.cells[util.PlaneIndex(w.Configs.Col, i, util.RightDown)].IsLive {
			w.cells[i].Score += 1
		}
		if w.cells[util.PlaneIndex(w.Configs.Col, i, util.Down)].IsLive {
			w.cells[i].Score += 1
		}
		if w.cells[util.PlaneIndex(w.Configs.Col, i, util.LeftDown)].IsLive {
			w.cells[i].Score += 1
		}
		if w.cells[util.PlaneIndex(w.Configs.Col, i, util.Left)].IsLive {
			w.cells[i].Score += 1
		}
		if w.cells[util.PlaneIndex(w.Configs.Col, i, util.LeftUp)].IsLive {
			w.cells[i].Score += 1
		}
	}

	return w
}

func (w World) EvalScore() World {
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

func (w World) ResetScore() World {
	for i, _ := range w.cells {
		w.cells[i].Score = 0
	}
	return w
}

func (w World) Draw() {
	for i, c := range w.cells {
		fmt.Print(c.String())
		// 改行
		if (i+1)%int(w.Configs.Col) == 0 {
			fmt.Println("")
		}
	}
}
