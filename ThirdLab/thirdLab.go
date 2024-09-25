package main

import (
	"GO/ThirdLab/mathutils"
	"GO/ThirdLab/stringutils"
	"fmt"
)

func findLongestString(slice []string) string {
	if len(slice) == 0 {
		return ""
	}

	longest := slice[0]
	for _, str := range slice {
		if len(str) > len(longest) {
			longest = str
		}
	}
	return longest
}

func main() {
	//Задание 1 & 2
	fmt.Println("Task one")
	var a int
	fmt.Print("Enter value for factorial: ")
	fmt.Scanln(&a)
	fmt.Println("Factorial: ", mathutils.Factorial(a))

	//Задание 3
	fmt.Println("\nTask two")
	myString := "12345"
	fmt.Println("Reverse: ", stringutils.Reverse(myString))

	//Задание 4
	fmt.Println("\nTask three")
	var myArray [5]int
	for i := 0; i < len(myArray); i++ {
		myArray[i] = i
	}
	fmt.Println(myArray)

	//Задание 5
	fmt.Println("\nTask four")
	slice := myArray[1:4]
	fmt.Println("Slice: ", slice)

	slice = append(slice, 6)
	fmt.Println("Slice after append: ", slice)

	slice = append(slice[:1], slice[2:]...)
	fmt.Println("Slice after pop:", slice)

	//Задание 6
	fmt.Println("\nTask six")
	strings := []string{"Go", "Gopher", "Программирование", "Самая длинная строка"}
	longest := findLongestString(strings)
	fmt.Println("Longest string: ", longest)
}
