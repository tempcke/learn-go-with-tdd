package bowling_game

type Game struct {
	rolls       [21]int
	currentRoll int
}

func (g *Game) Roll(pins int) {
	g.rolls[g.currentRoll] = pins
	g.currentRoll++
}

func (g *Game) Score() int {
	score := 0
	i := 0

	for frame := 0; frame < 10; frame++ {
		ball1 := g.rolls[i]
		ball2 := g.rolls[i+1]
		if ball1 == 10 {
			score += 10 + g.rolls[i+1] + g.rolls[i+2]
			i += 1
		} else if ball1+ball2 == 10 {
			score += 10 + g.rolls[i+2]
			i += 2
		} else {
			score += ball1 + ball2
			i += 2
		}
	}

	return score
}
