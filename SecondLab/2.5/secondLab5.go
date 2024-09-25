package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	//Задание 5
	fmt.Println("\nTask five.")
	rect := Rectangle{Width: 5, Height: 10}
	fmt.Println("Area: ", rect.Area())
}
