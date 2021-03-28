package models

import "go-snake/src/ui"

const (
	Up Direction = iota
	Down
	Left
	Right
)
const snakeSize = 5

var snakeInitialHeadPosition = ui.Pixel{
	X: 10,
	Y: 20,
}

type Direction int

type Snake struct {
	head ui.Pixel
	body []ui.Pixel
	direction Direction
}

func (s *Snake) Draw()  {
	for _, pixel := range s.body {
		ui.Write(pixel, "*")
	}
}

func (s *Snake) Init()  {
	pixel := snakeInitialHeadPosition

	s.head = pixel

	i := 0
	for i < snakeSize {
		s.body = append(s.body, pixel)
		pixel.X += 1
		i++
	}
}

func (s *Snake) Move() {
	s.head.Y -= 1
	s.body = s.body[:len(s.body) - 1]
	s.body = append([]ui.Pixel{s.head}, s.body...)
}


