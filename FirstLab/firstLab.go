package main

import (
	"fmt"
	"time"
)

func doSomeIntOperation() {
	var a, b int
	var operation string

	fmt.Println("First value: ")
	fmt.Scan(&a)
	fmt.Println("\nSecond value: ")
	fmt.Scan(&b)
	fmt.Println("\nOperation (+, -, *, /): ")
	fmt.Scan(&operation)

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

func doSomeFloatOperation(a, b float32) (float32, float32) {
	return a + b, a - b
}

func countAverageNumber(a, b, c int) float64 {
	result := float64(a+b+c) / 3
	return result
}

func main() {
	// Задание 1
	fmt.Println("Task one.\nTime: ", time.Now())

	// Задание 2
	var intValue int = 64
	var floatValue float64 = 777.777
	var stringValue string = "stringValue"
	var boolValue bool = true
	fmt.Println("\nTask Two.\nInt: ", intValue, "\nfloatValue: ", floatValue, "\nstringValue: ", stringValue, "\nboolValue: ", boolValue)

	//Задание 3
	secondIntValue := 128
	fmt.Println("\nTask Three\n", secondIntValue)

	//Задание 4
	fmt.Println("\nTask Four")
	doSomeIntOperation()

	//Задание 5
	fmt.Println("\nTask Five")
	fmt.Println(doSomeFloatOperation(22.2, 7.8))

	//Задание 6
	fmt.Println("\nTask Six")
	fmt.Println(countAverageNumber(10, 20, 30))
}
