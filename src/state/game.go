package state

const (
	BoardDimension Item = iota // the dimensions of the snake board (height and width)
	CurrentRound // the current round (each move of the snake constitutes a round)
	Score // Score (number of food units eaten)
	SnakeLength // length of the snake
	SnakeHeadCoordinate // coordinates of snake head
)
type Item int

type GameState [][]int

func save()  {
}

