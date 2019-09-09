package bowling_game


type Game struct {
	roles []int
}

func (g *Game) Roll(pins int) {
	g.roles = append(g.roles, pins)
}

func (g *Game) Score() int {
	sum := 0

	for i := range g.roles {
		sum += g.roles[i]
    }

	return sum
}