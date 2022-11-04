package main

import (
	"bufio"
	"fmt"
	"golife/world"
	"os"
)

func main() {
	w := world.LoadWorld()

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
