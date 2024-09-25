package main

import "fmt"

func main() {
	fmt.Println("\nTask Four")
	var a, b int
	var operation string

	fmt.Print("First value: ")
	fmt.Scanln(&a)
	fmt.Print("\nSecond value: ")
	fmt.Scanln(&b)
	fmt.Print("\nOperation (+, -, *, /): ")
	fmt.Scanln(&operation)

	switch operation {
	case "+":
		fmt.Println("Result: ", a+b)
	case "-":
		fmt.Println("Result: ", a-b)
	case "*":
		fmt.Println("Result: ", a*b)
	case "/":
		if b != 0 {
			fmt.Println("Result: ", a/b)
		} else {
			fmt.Println("Error: / 0!")
		}
	default:
		fmt.Println("Error: Unsupported operation")
	}
}
