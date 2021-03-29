package action

import (
	"fmt"
)

var input string

const (
	NONE Direction = iota
	UP
	DOWN
	LEFT
	RIGHT
)

type Direction int

func Ask() (Direction, error) {

	for {
		fmt.Print("Enter direction [WASD for UP, LEFT, DOWN, RIGHT respectively] and press ENTER: ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			return NONE, err
		}

		direction := GetDirection(input)
		if direction != NONE {
			return direction, nil
		}
	}
}

func GetDirection(input string) Direction {
	switch input {
	case "w", "W":
		return UP
	case "a", "A":
		return LEFT
	case "s", "S":
		return DOWN
	case "d", "D":
		return RIGHT
	}

	return NONE
}
