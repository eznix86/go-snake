package ui

import "fmt"
// Using
// To draw on ui
// https://misc.flogisoft.com/bash/tip_colors_and_formatting
// https://tldp.org/HOWTO/Bash-Prompt-HOWTO/x361.html
// A Pixel represent a coordinate of on the UI
type Pixel struct {
	X int
	Y int
}

// Helps to Write a text to the screen at a defined Pixel of the terminal
func Write(pixel Pixel, text string)  {
	fmt.Printf("\033[%d;%dH%s", pixel.Y, pixel.X*3, text)
}

func Clear()  {
	fmt.Println("\033[2J")
}