package main

import (
	"bufio"
	"golife/world"
	"os"
)

func main() {
	h := &world.History{}
	w := world.LoadWorld()
	h.Worlds = h.CreateHistory(*w)

	i := 0
	h.Worlds[i].Draw()
	i += 1

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		h.Worlds[i%5].Draw()
		i += 1
	}
}
