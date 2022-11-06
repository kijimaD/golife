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
	Cells   []Cell
	configs config.Configs
}

// TODO: 今はファイルだけ。Readerで読み込めるようにする
// 外部入力から初期世界を生成する
func Load(c config.Configs) *World {
	w := &World{
		configs: c,
	}

	data, _ := ioutil.ReadFile(INPUT_FILE)
	for _, rune := range string(data) {
		switch string(rune) {
		case LIVEC:
			w.Cells = append(w.Cells, NewCell(true))
		case DEATHC:
			w.Cells = append(w.Cells, NewCell(false))
		case "\n":
		default:
			fmt.Printf("`%s`は不正な文字です\n", string(rune))
		}
	}

	return w
}

func (w World) Next() World {
	return w.resetScore().calcScore().evalScore()
}

func (w World) calcScore() World {
	for i, _ := range w.Cells {
		if w.Cells[util.PlaneIndex(w.configs.Col, i, util.Up)].IsLive {
			w.Cells[i].Score += 1
		}
		if w.Cells[util.PlaneIndex(w.configs.Col, i, util.RightUp)].IsLive {
			w.Cells[i].Score += 1
		}
		if w.Cells[util.PlaneIndex(w.configs.Col, i, util.Right)].IsLive {
			w.Cells[i].Score += 1
		}
		if w.Cells[util.PlaneIndex(w.configs.Col, i, util.RightDown)].IsLive {
			w.Cells[i].Score += 1
		}
		if w.Cells[util.PlaneIndex(w.configs.Col, i, util.Down)].IsLive {
			w.Cells[i].Score += 1
		}
		if w.Cells[util.PlaneIndex(w.configs.Col, i, util.LeftDown)].IsLive {
			w.Cells[i].Score += 1
		}
		if w.Cells[util.PlaneIndex(w.configs.Col, i, util.Left)].IsLive {
			w.Cells[i].Score += 1
		}
		if w.Cells[util.PlaneIndex(w.configs.Col, i, util.LeftUp)].IsLive {
			w.Cells[i].Score += 1
		}
	}

	return w
}

func (w World) evalScore() World {
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

func (w World) resetScore() World {
	for i, _ := range w.Cells {
		w.Cells[i].Score = 0
	}
	return w
}

func (w World) Draw() {
	for i, c := range w.Cells {
		fmt.Print(c.String())
		// 改行
		if (i+1)%int(w.configs.Col) == 0 {
			fmt.Println("")
		}
	}
}
