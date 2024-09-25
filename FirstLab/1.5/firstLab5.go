package main

import "fmt"

func doSomeFloatOperation(a, b float32) (float32, float32) {
	return a + b, a - b
}

func main() {
	//Задание 5
	fmt.Println("\nTask Five")
	sum, minus := doSomeFloatOperation(22.2, 7.8)
	fmt.Println("Sum: ", sum, "\nMin: ", minus)
}
