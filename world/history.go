package world

import (
	"golife/config"
)

type History struct {
	Worlds []World
}

func (h History) CreateHistory(initialWorld World, c config.Configs) []World {
	var worlds []World
	var w World
	w = initialWorld

	// 初期状態
	// コピーしてリストに入れる
	cells := make([]Cell, len(w.Cells))
	copy(cells, w.Cells)
	cp := World{
		Cells:   cells,
		configs: c,
	}
	worlds = append(worlds, cp)

	for i := 0; i < c.GenCap; i++ {
		w = w.Next()
		cells := make([]Cell, len(w.Cells))
		copy(cells, w.Cells)
		cp := World{
			Cells:   cells,
			configs: c,
		}
		worlds = append(worlds, cp)
	}

	return worlds
}
