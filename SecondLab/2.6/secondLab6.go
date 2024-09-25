package main

import "fmt"

func countAverageNumber(a, b int) float64 {
	return float64((a + b) / 2)
}

func main() {
	//Задание 6
	fmt.Println("\nTask six.")
	fmt.Println("Average: ", countAverageNumber(10, 20))
}
