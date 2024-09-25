package main

import "fmt"

func main() {
	var myArray [5]int
	//Задание 5
	fmt.Println("\nTask four")
	slice := myArray[1:4]
	fmt.Println("Slice: ", slice)

	slice = append(slice, 6)
	fmt.Println("Slice after append: ", slice)

	slice = append(slice[:1], slice[2:]...)
	fmt.Println("Slice after pop:", slice)
}
