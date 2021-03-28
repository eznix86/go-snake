package main

import (
	"go-snake/src/models"
	"time"
)

func main() {
	food := models.Food{}
	snake := models.Snake{}
	snake.Init()
	food.RandomizeLocation()

	for {
		canvas := models.NewBoard()
		canvas.Draw()
		food.Draw()
		snake.Draw()
		snake.Move()
		time.Sleep(time.Duration(500) * time.Millisecond)
	}
}