package world

const (
	LIVEC  = `●`
	DEATHC = `○`
)

type Cell struct {
	IsLive bool
	score  int
}

func (c Cell) String() string {
	if c.IsLive {
		return LIVEC
	} else {
		return DEATHC
	}
}

func NewCell(isLive bool) Cell {
	return Cell{isLive, 0}
}
