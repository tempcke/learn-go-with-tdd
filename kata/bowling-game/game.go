package bowling_game

type Game struct {
	rolls [21]int
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
		frameScore := g.rolls[i] + g.rolls[i+1]
		if frameScore == 10 {
			score += g.rolls[i+2]
		}
		score += frameScore
		i += 2
	}

	return score
}
