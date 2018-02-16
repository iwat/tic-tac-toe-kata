package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTicTacToe3XWin(t *testing.T) {
	game := ticTacToe{}
	game.init(3, xPlayer)
	dump(t, game)

	winner := nilPlayer
	err := error(nil)

	winner, err = game.mark(0)
	dump(t, game)
	assert.Equal(t, nilPlayer, winner)
	assert.Nil(t, err)

	winner, err = game.mark(1)
	dump(t, game)
	assert.Equal(t, nilPlayer, winner)
	assert.Nil(t, err)

	winner, err = game.mark(4)
	dump(t, game)
	assert.Equal(t, nilPlayer, winner)
	assert.Nil(t, err)

	winner, err = game.mark(2)
	dump(t, game)
	assert.Equal(t, nilPlayer, winner)
	assert.Nil(t, err)

	winner, err = game.mark(8)
	dump(t, game)
	assert.Equal(t, xPlayer, winner)
	assert.Nil(t, err)

	_, err = game.mark(3)
	dump(t, game)
	assert.Equal(t, gameFinishedErr, err)
}

func TestTicTacToe3Draw(t *testing.T) {
	game := ticTacToe{}
	game.init(3, xPlayer)
	dump(t, game)

	winner := nilPlayer
	err := error(nil)

	winner, err = game.mark(0)
	dump(t, game)
	assert.Equal(t, nilPlayer, winner)
	assert.Nil(t, err)

	winner, err = game.mark(4)
	dump(t, game)
	assert.Equal(t, nilPlayer, winner)
	assert.Nil(t, err)

	winner, err = game.mark(1)
	dump(t, game)
	assert.Equal(t, nilPlayer, winner)
	assert.Nil(t, err)

	winner, err = game.mark(2)
	dump(t, game)
	assert.Equal(t, nilPlayer, winner)
	assert.Nil(t, err)

	winner, err = game.mark(6)
	dump(t, game)
	assert.Equal(t, nilPlayer, winner)
	assert.Nil(t, err)

	winner, err = game.mark(3)
	dump(t, game)
	assert.Equal(t, nilPlayer, winner)
	assert.Nil(t, err)

	winner, err = game.mark(5)
	dump(t, game)
	assert.Equal(t, nilPlayer, winner)
	assert.Nil(t, err)

	winner, err = game.mark(7)
	dump(t, game)
	assert.Equal(t, nilPlayer, winner)
	assert.Nil(t, err)

	winner, err = game.mark(8)
	dump(t, game)
	assert.Equal(t, nilPlayer, winner)
	assert.Nil(t, err)

	_, err = game.mark(3)
	dump(t, game)
	assert.Equal(t, gameFinishedErr, err)
}

func TestTicTacToe5XWin(t *testing.T) {
	game := ticTacToe{}
	game.init(5, xPlayer)
	dump(t, game)

	winner := nilPlayer
	err := error(nil)

	winner, err = game.mark(0)
	dump(t, game)
	assert.Equal(t, nilPlayer, winner)
	assert.Nil(t, err)

	winner, err = game.mark(1)
	dump(t, game)
	assert.Equal(t, nilPlayer, winner)
	assert.Nil(t, err)

	winner, err = game.mark(6)
	dump(t, game)
	assert.Equal(t, nilPlayer, winner)
	assert.Nil(t, err)

	winner, err = game.mark(2)
	dump(t, game)
	assert.Equal(t, nilPlayer, winner)
	assert.Nil(t, err)

	winner, err = game.mark(12)
	dump(t, game)
	assert.Equal(t, nilPlayer, winner)
	assert.Nil(t, err)

	winner, err = game.mark(3)
	dump(t, game)
	assert.Equal(t, nilPlayer, winner)
	assert.Nil(t, err)

	winner, err = game.mark(18)
	dump(t, game)
	assert.Equal(t, nilPlayer, winner)
	assert.Nil(t, err)

	winner, err = game.mark(4)
	dump(t, game)
	assert.Equal(t, nilPlayer, winner)
	assert.Nil(t, err)

	winner, err = game.mark(24)
	dump(t, game)
	assert.Equal(t, xPlayer, winner)
	assert.Nil(t, err)

	_, err = game.mark(3)
	dump(t, game)
	assert.Equal(t, gameFinishedErr, err)
}

func dump(t *testing.T, game ticTacToe) {
	for y := 0; y < game.size; y++ {
		line := &bytes.Buffer{}
		for x := 0; x < game.size; x++ {
			fmt.Fprintf(line, "%d ", game.state[x+game.size*y])
		}
		t.Log(line.String())
	}
	t.Log("---------")
}
