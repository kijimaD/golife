package world

import (
	"golife/config"
)

type History struct {
	Worlds  []World
	Configs config.Configs
}

func (h History) CreateHistory(initialWorld World) []World {
	var worlds []World
	var w World
	w = initialWorld

	// 初期状態
	// コピーしてリストに入れる
	cells := make([]Cell, len(w.Cells))
	copy(cells, w.Cells)
	cp := World{
		Cells:   cells,
		configs: h.Configs,
	}
	worlds = append(worlds, cp)

	for i := 0; i < h.Configs.GenCap; i++ {
		w = w.Next()
		cells := make([]Cell, len(w.Cells))
		copy(cells, w.Cells)
		cp := World{
			Cells:   cells,
			configs: h.Configs,
		}
		worlds = append(worlds, cp)
	}

	return worlds
}
