package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/theArtechnology/go-snake/src/action"
	"github.com/theArtechnology/go-snake/src/game"
)

var (
	err       error
	boardSize = 10
)

const ErrorDisplayTime = 3

func main() {
	flag.IntVar(&boardSize, "size", 10, "Size of the board")
	flag.Parse()
	if boardSize <= 5 {
		fmt.Println("Board too small for the snake! Give a bigger size !")
		os.Exit(0)
	}
	board := game.New(boardSize, boardSize)
	board.Init()

	for {
		err = board.MoveSnake()
		board.FindFood()
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second * time.Duration(ErrorDisplayTime))
			break
		}
		board.Draw()
		fmt.Println()
		fmt.Printf("Current Round: %d", board.CurrentRound)
		fmt.Println()
		fmt.Printf("Score: %d", board.Score)
		fmt.Println()
		currentDirection, err := action.Ask()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Unknown user action: retrying...")
			time.Sleep(time.Second * time.Duration(ErrorDisplayTime))
		}

		err = board.SetDirection(currentDirection)
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second * time.Duration(ErrorDisplayTime))
		}
	}
}
