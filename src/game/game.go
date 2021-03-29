package game

import (
	"fmt"
	"go-snake/src/action"
	"math/rand"
	"time"
)

// tiles
const (
	Empty  Cell = iota
	Fences      // aka boundary
	Snake
	Food
)

type Cell int

type Board struct {
	Width  int
	Height int
	// gamestate will not be an "array" ([x][y]int) per se, as it will grow dynamically with the board size, hence use of slice.
	// It will  use a stride to create the slice in question.
	gamestate []int
	// Used for offset in the Y coordinates in the gamestate
	stride       int
	Score        int
	CurrentRound int
	SnakeLength  int
	SnakeHead    int
	SnakeBody    []int
	FoodLocation int
	direction    action.Direction
}

func New(width int, height int) *Board {
	board := &Board{
		Width:  width,
		Height: height,
	}
	// use to create a slice
	board.gamestate = make([]int, width*height)
	board.stride = width
	board.direction = action.NONE
	return board
}

func (b *Board) calculateCell(x, y int) int {
	return y*b.stride + x
}

func (b *Board) SetDirection(direction action.Direction) error {
	switch direction {
	case action.UP:
		if b.direction == action.DOWN {
			return fmt.Errorf("you can't do a 180 turn")
		}
	case action.DOWN:
		if b.direction == action.UP {
			return fmt.Errorf("you can't do a 180 turn")
		}
	case action.LEFT, action.NONE:
		if b.direction == action.RIGHT {
			return fmt.Errorf("you can't do a 180 turn")
		}
	case action.RIGHT:
		if b.direction == action.LEFT {
			return fmt.Errorf("you can't do a 180 turn")
		}
	}

	b.direction = direction
	return nil
}

func (b *Board) SetItemAtPosition(position int, value Cell) {
	// equivalent to gamestate[x][y] = value,
	// where stride is the Y coordinate
	b.gamestate[position] = int(value)
}

func (b *Board) SetItem(x int, y int, value Cell) {
	// equivalent to gamestate[x][y] = value,
	// where stride is the Y coordinate
	b.SetItemAtPosition(b.calculateCell(x, y), value)
}

func (b *Board) GetItem(x int, y int) Cell {
	return b.GetItemAtPosition(b.calculateCell(x, y))
}

func (b *Board) GetItemAtPosition(position int) Cell {
	return Cell(b.gamestate[position])
}

func (b *Board) buildFence() {
	x := 0
	y := 0
	for x < b.Width {
		b.SetItem(x, y, Fences)
		x += 1
	}

	x -= 1
	for y < b.Height {
		b.SetItem(x, y, Fences)
		y += 1
	}

	y -= 1
	for x > 0 {
		b.SetItem(x, y, Fences)
		x -= 1
	}

	for y > 0 {
		b.SetItem(x, y, Fences)
		y -= 1
	}
}

func (b *Board) buildSnakeBody() {
	b.SnakeLength = 3
	// centered head
	x := b.stride / 2
	y := b.stride / 2

	b.SnakeHead = b.calculateCell(x, y)

	body := 0
	for body < b.SnakeLength {
		// build from the head to the rest of the body
		b.SnakeBody = append(b.SnakeBody, b.calculateCell(x, y-body))
		body += 1
	}
	// add them into gamestate
	b.drawSnake()
}

// Good ol' snack
func (b *Board) generateFood() {
	rand.Seed(time.Now().UTC().UnixNano())
	min := 0
	max := b.Width * b.Height

	for {
		// add the food the a random position where it is empty
		position := rand.Intn(max-min) + min
		if b.GetItemAtPosition(position) == Empty {
			if b.FoodLocation != 0 {
				previousLocation := b.FoodLocation
				b.SetItemAtPosition(previousLocation, Empty)
			}
			b.FoodLocation = position
			b.SetItemAtPosition(position, Food)
			break
		}
	}
}

func (b *Board) GetXYFromPosition(position int) (int, int) {
	x := position % b.stride
	y := (position) / b.stride
	return x, y
}

func (b *Board) Init() {
	b.buildFence()
	b.buildSnakeBody()
	b.generateFood()
}

func (b *Board) MoveSnake() error {

	if b.direction == action.NONE {
		return nil
	}

	x, y := b.GetXYFromPosition(b.SnakeHead)

	headLocation := b.calculateCell(x, y)
	switch b.direction {
	case action.UP:
		headLocation = b.calculateCell(x-1, y)

	case action.DOWN:
		headLocation = b.calculateCell(x+1, y)

	case action.LEFT:
		headLocation = b.calculateCell(x, y-1)

	case action.RIGHT:
		headLocation = b.calculateCell(x, y+1)
	}

	// skip own head
	RestOfTheBodyOfTheSnake := b.SnakeBody[1:]
	for _, body := range RestOfTheBodyOfTheSnake {
		if headLocation == body {
			return fmt.Errorf("you've bite your own tail")
		}
	}

	if b.GetItemAtPosition(headLocation) == Fences {
		return fmt.Errorf("your head hit on a fence")
	}

	for _, snake := range b.SnakeBody {
		b.SetItemAtPosition(snake, Empty)
	}

	b.SnakeBody = append([]int{headLocation}, b.SnakeBody[:b.SnakeLength-1]...)
	b.SnakeHead = headLocation
	b.CurrentRound += 1
	b.drawSnake()

	return nil

}

func (b *Board) drawSnake() {
	for _, snake := range b.SnakeBody {
		b.SetItemAtPosition(snake, Snake)
	}
}

func (b *Board) FindFood() {
	if b.FoodLocation == b.SnakeHead {
		b.Score += 1
		b.generateFood()
		b.drawSnake()
		b.SnakeLength += 1
	}
}

func (b *Board) Draw() {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorBlue := "\033[34m"

	for x := 0; x < b.stride; x++ {
		fmt.Println()
		for y := 0; y < b.stride; y++ {
			switch b.GetItem(x, y) {
			case Empty:
				fmt.Print(" ")
				break
			case Snake:
				fmt.Print(string(colorGreen), "#", string(colorReset))
				break
			case Fences:
				fmt.Print(string(colorRed), "+", string(colorReset))
				break
			case Food:
				fmt.Print(string(colorBlue), "*", string(colorReset))
				break
			}
			fmt.Print("  ")
		}
	}

	fmt.Println()
}
