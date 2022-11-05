package config

import (
	"bufio"
	"io"
	"os"
)

const INPUT_FILE = "world.txt"

type Configs struct {
	Debug bool
	Row   int
	Col   int
}

func Load() Configs {
	var row int
	var col int

	f, _ := os.Open(INPUT_FILE)
	bu := bufio.NewReaderSize(f, 1024)
	for {
		line, _, err := bu.ReadLine()
		if err == io.EOF {
			break
		}
		// 最終行の文字長をcolとする
		col = len([]rune(string(line)))
		// 行数をrowとする
		row += 1
	}
	f.Close()

	return Configs{
		Debug: false,
		Row:   row,
		Col:   col,
	}
}
