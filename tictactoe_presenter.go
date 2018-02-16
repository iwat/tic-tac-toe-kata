package main

import (
	"strings"
)

type ticTacToePresenter struct {
	ticTacToe
}

func (p ticTacToePresenter) present() string {
	lines := make([]string, p.size)
	for y := 0; y < p.size; y++ {
		row := make([]string, p.size)
		for x := 0; x < p.size; x++ {
			switch p.state[x+p.size*y] {
			case xPlayer:
				row[x] = "X"
			case oPlayer:
				row[x] = "O"
			}
		}
		line := strings.Join(row, " | ")
		lines[y] = line
	}
	return strings.Join(lines, "\n")
}
