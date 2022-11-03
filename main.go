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
