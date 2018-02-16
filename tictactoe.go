package main

import "errors"

var (
	blockOccupiedErr   = errors.New("block is occupied")
	blockOutOfRangeErr = errors.New("block number is out of range")
	gameFinishedErr    = errors.New("game is finished")
)

type player byte

const (
	nilPlayer player = 0
	xPlayer   player = 1
	oPlayer   player = 2
)

func (p player) next() player {
	switch p {
	case xPlayer:
		return oPlayer
	case oPlayer:
		return xPlayer
	default:
		return nilPlayer
	}
}

type ticTacToe struct {
	size     int
	state    []player
	turn     player
	finished bool
}

func (t *ticTacToe) init(size int, first player) {
	t.size = size
	t.state = make([]player, size*size)
	t.turn = first
	t.finished = false
}

func (t *ticTacToe) mark(block int) (player, error) {
	if t.finished {
		return nilPlayer, gameFinishedErr
	}
	if block < 0 || block > t.size*t.size {
		return nilPlayer, blockOutOfRangeErr
	}
	if t.state[block] != nilPlayer {
		return nilPlayer, blockOccupiedErr
	}
	t.state[block] = t.turn
	if t._finished() {
		t.finished = true
		return t.turn, nil
	}
	t.turn = t.turn.next()
	return nilPlayer, nil
}

func (t ticTacToe) _finished() bool {
	idx := []int(nil)
	for i := 0; i < t.size; i++ {
		idx = append(idx, i+t.size*i)
	}
	if t._check(idx) {
		return true
	}

	idx = nil
	for i := t.size - 1; i >= 0; i-- {
		idx = append(idx, i+t.size*i)
	}
	if t._check(idx) {
		return true
	}

	idx = nil
	for x := 0; x < t.size; x++ {
		for y := 0; y < t.size; y++ {
			idx = append(idx, x+t.size*y)
		}
	}
	if t._check(idx) {
		return true
	}

	for y := 0; y < t.size; y++ {
		for x := 0; x < t.size; x++ {
			idx = append(idx, x+t.size*y)
		}
	}
	if t._check(idx) {
		return true
	}

	return false
}

func (t ticTacToe) _check(idx []int) bool {
	// since we only call this after marking, the winner can only be the last turn.
	for _, i := range idx {
		if t.state[i] != t.turn {
			return false
		}
	}
	return true
}
