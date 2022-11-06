// ライフゲームをターミナルで起動する

package cli

import (
	"bufio"
	"fmt"
	"golife/config"
	"golife/world"
	"os"
)

const (
	NEXT_KEY = "n"
	PREV_KEY = "p"
)

func Run() {
	h := &world.History{}
	c := config.Load()
	w := world.Load(c)
	h.Worlds = h.CreateHistory(*w, c)

	fmt.Print("[n]ext or [p]rev\n")
	i := 0
	fmt.Printf("gen: %d\n", i)
	h.Worlds[i].Draw()

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "", NEXT_KEY:
			i += 1
		case PREV_KEY:
			i -= 1
		default:
			continue
		}

		fmt.Printf("gen: %d\n", i)
		h.Worlds[i%c.GenCap].Draw()
	}
}
