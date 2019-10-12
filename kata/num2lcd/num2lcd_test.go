package num2lcd

import (
	"gotest.tools/assert"
	"testing"
)

func TestZero(t *testing.T) {
	expected := " _ \n| |\n|_|\n"
	assert.Equal(t, expected, Num2Lcd(0))
}

func TestOne(t *testing.T) {
	expected := "   \n  |\n  |\n"
	assert.Equal(t, expected, Num2Lcd(1))
}

func TestTwo(t *testing.T) {
	expected := " _ \n _|\n|_ \n"
	assert.Equal(t, expected, Num2Lcd(2))
}

func TestThree(t *testing.T) {
	expected := " _ \n _|\n _|\n"
	assert.Equal(t, expected, Num2Lcd(3))
}

func TestFour(t *testing.T) {
	expected := "   \n|_|\n  |\n"
	assert.Equal(t, expected, Num2Lcd(4))
}

func TestFive(t *testing.T) {
	expected := " _ \n|_ \n _|\n"
	assert.Equal(t, expected, Num2Lcd(5))
}

func TestSix(t *testing.T) {
	expected := " _ \n|_ \n|_|\n"
	assert.Equal(t, expected, Num2Lcd(6))
}

func TestSeven(t *testing.T) {
	expected := " _ \n  |\n  |\n"
	assert.Equal(t, expected, Num2Lcd(7))
}

func TestEight(t *testing.T) {
	expected := " _ \n|_|\n|_|\n"
	assert.Equal(t, expected, Num2Lcd(8))
}

func TestNine(t *testing.T) {
	expected := " _ \n|_|\n _|\n"
	assert.Equal(t, expected, Num2Lcd(9))
}
