package main

import "fmt"

func main() {
	//Задание 6
	fmt.Println("\nTask Six")
	var a, b, c int
	fmt.Print("First value: ")
	fmt.Scanln(&a)
	fmt.Print("\nSecond value: ")
	fmt.Scanln(&b)
	fmt.Print("Third value: ")
	fmt.Scanln(&c)

	fmt.Println("Average: ", float64(a+b+c)/3)
}
