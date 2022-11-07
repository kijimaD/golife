package config

import (
	"bufio"
	"github.com/go-yaml/yaml"
	"github.com/labstack/echo/v4"
	"io"
	"os"
	"strconv"
)

const INPUT_FILE = "world.txt"
const CONFIG_FILE = ".golife.yml"

type Configs struct {
	Debug  bool `yaml:"Debug"`
	GenCap int  `yaml:"GenCap"`
	Row    int
	Col    int
}

func new() *Configs {
	return &Configs{
		Debug:  false,
		GenCap: 100,
		Row:    100,
		Col:    100,
	}
}

// CLIモードでは設定ファイル(.yml)と入力ファイル(.txt)からConfigsを生成する。
func CLILoad() Configs {
	// configsをここでinitして、各メソッドで読み込む方式にする
	c := new()

	f, _ := os.Open(INPUT_FILE)
	defer f.Close()
	bu := bufio.NewReaderSize(f, 1024)
	c.loadSize(bu).loadConfigYML()

	return *c
}

// APIモードではリクエストからConfigsを生成する。
func ServerLoad(con echo.Context) Configs {
	c := new()

	// 暫定でファイルから
	f, _ := os.Open(INPUT_FILE)
	defer f.Close()
	bu := bufio.NewReaderSize(f, 1024)

	c.loadSize(bu)
	c.GenCap, _ = strconv.Atoi(con.FormValue("GenCap"))
	c.Debug = con.FormValue("Debug") == "true"

	// c.loadReq()

	return *c
}

// row, colをセットする
func (c *Configs) loadSize(r io.Reader) *Configs {
	var row int
	var col int

	bu := bufio.NewReaderSize(r, 1024)
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

	c.Row = row
	c.Col = col
	return c
}

// Debug, Gencapをセットする
func (c *Configs) loadConfigYML() *Configs {
	f, err := os.Open(CONFIG_FILE)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(c)
	if err != nil {
		panic(err)
	}

	return c
}
