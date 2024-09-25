package main

import "fmt"

func main() {
	var a int

	//Задание 1
	fmt.Println("Task one.")
	fmt.Print("Enver value: ")
	fmt.Scan(&a)

	if a%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("not even")
	}
}
