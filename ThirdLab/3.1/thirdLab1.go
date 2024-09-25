package main

import (
	"GO/ThirdLab/mathutils"
	"fmt"
)

func main() {
	//Задание 1
	fmt.Println("Task one")
	var a int
	fmt.Print("Enter value for factorial: ")
	fmt.Scanln(&a)
	fmt.Println("Factorial: ", mathutils.Factorial(a))
}
