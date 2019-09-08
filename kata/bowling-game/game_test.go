package bowling_game

import (
	"gotest.tools/assert"
	"testing"
)

func TestGutterGame(t *testing.T) {
  g := Game{}
  for i:=0; i<20; i++ {
  	g.Roll(0)
  }
  assert.Equal(t, 0, g.Score())
}
