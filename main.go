package main

import (
	"bufio"
	"fmt"
	"golife/world"
	"os"
)

func main() {
	w := &world.World{
		Row: 5,
		Col: 5,
		Cells: []world.Cell{
			world.NewCell(false),
			world.NewCell(false),
			world.NewCell(false),
			world.NewCell(false),
			world.NewCell(false),

			world.NewCell(false),
			world.NewCell(true),
			world.NewCell(true),
			world.NewCell(true),
			world.NewCell(false),

			world.NewCell(false),
			world.NewCell(false),
			world.NewCell(true),
			world.NewCell(true),
			world.NewCell(false),

			world.NewCell(false),
			world.NewCell(false),
			world.NewCell(false),
			world.NewCell(false),
			world.NewCell(false),

			world.NewCell(false),
			world.NewCell(false),
			world.NewCell(false),
			world.NewCell(false),
			world.NewCell(false),
		},
	}

	w.Draw()
	gen := 0

	for true {
		fmt.Printf("%d gen: enter\n", gen)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		w.ResetScore().CalcScore().EvalScore().Draw()
		gen += 1
	}
}
