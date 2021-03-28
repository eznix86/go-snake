package models

import (
	"go-snake/src/state"
	"go-snake/src/ui"
)

// A board represent the boundary of the state
// This will depend of the terminal size
type Board struct {
	height int
	width int
	state *state.GameState
}

func NewBoard() *Board {
	return &Board{
		height: 20,
		width: 20,
	}
}

func (b *Board) Set(state *state.GameState)  {
	b.state = state
}

func (b *Board) Save()  {

}

func (b *Board) Draw() ui.Pixel {
	n := 0
	ui.Clear()

	pixel := ui.Pixel{
		X: 5,
		Y: 5,
	}
	for n < b.width {
		n += 1
		pixel.X += 1
		ui.Write(pixel, "—")
	}

	n = 0
	for n < b.height {
		n += 1
		pixel.Y += 1
		ui.Write(pixel, "|")
	}

	n = 0
	for n < b.height {
		n += 1
		pixel.X -= 1
		ui.Write(pixel, "_")
	}

	n = 0
	for n < b.height {
		n += 1
		pixel.Y -= 1
		ui.Write(pixel, "|")
	}

	return pixel
}