package bowling_game


type Game struct {
	roles []int
}

func (g Game) Roll(pins int) {
	g.roles = append(g.roles, pins)
}

func (g Game) Score() int {
	return 0
}