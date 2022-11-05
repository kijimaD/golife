package config

type Configs struct {
	Debug bool
	Row   int
	Col   int
}

func New() *Configs {
	return &Configs{
		Debug: false,
		Row:   0,
		Col:   0,
	}
}
