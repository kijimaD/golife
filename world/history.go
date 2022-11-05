package world

type History struct {
	Worlds []World
}

func (h History) CreateHistory(initialWorld World) []World {
	var worlds []World
	var w World
	w = initialWorld

	// 初期状態
	// コピーしてリストに入れる
	cells := make([]Cell, len(w.cells))
	copy(cells, w.cells)
	cp := World{
		col:   w.col,
		row:   w.row,
		cells: cells,
	}
	worlds = append(worlds, cp)

	// とりあえず100ループ
	for i := 0; i < 100; i++ {
		w = w.Next()
		cells := make([]Cell, len(w.cells))
		copy(cells, w.cells)
		cp := World{
			col:   w.col,
			row:   w.row,
			cells: cells,
		}
		worlds = append(worlds, cp)
	}

	return worlds
}
