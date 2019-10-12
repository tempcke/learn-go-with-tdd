package num2lcd

import (
	"gotest.tools/assert"
	"testing"
)

type req struct {
	number int
	expected string
}

var reqs = []req{
	{0, " _ \n| |\n|_|\n"},
	{1, "   \n  |\n  |\n"},
	{2, " _ \n _|\n|_ \n"},
	{3, " _ \n _|\n _|\n"},
	{4, "   \n|_|\n  |\n"},
	{5, " _ \n|_ \n _|\n"},
	{6, " _ \n|_ \n|_|\n"},
	{7, " _ \n  |\n  |\n"},
	{8, " _ \n|_|\n|_|\n"},
	{9, " _ \n|_|\n _|\n"},
}

func TestSingleDigits(t *testing.T) {
	for _, req := range reqs {
		assert.Equal(t, req.expected, Num2Lcd(req.number))
	}
}