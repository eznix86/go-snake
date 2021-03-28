package models

import (
	"go-snake/src/ui"
	"math/rand"
	"time"
)

type Food struct {
	location ui.Pixel
}

func (f *Food) Draw()  {
	ui.Write(f.location, "#")
}

func (f *Food) RandomizeLocation()  {
	rand.Seed(time.Now().UTC().UnixNano())
	x := 10 + rand.Intn(30 - 10)
	y := 10 + rand.Intn(30 - 10)
	f.location = ui.Pixel{
		X: x,
		Y: y,
	}
}
