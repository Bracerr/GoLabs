package main

import "fmt"

func main() {
	//Задание 4
	fmt.Println("\nTask three")
	var myArray [5]int
	for i := 0; i < len(myArray); i++ {
		myArray[i] = i
	}
	fmt.Println(myArray)
}
