package main

import "fmt"

func determineNumber(number int) string {
	if number > 0 {
		return "Positive"
	} else if number < 0 {
		return "Negative"
	} else if number == 0 {
		return "Zero"
	}
	return "Error."
}

func main() {
	//Задание 2
	fmt.Println("\nTask two.")
	fmt.Println("10: ", determineNumber(10))
	fmt.Println("-11: ", determineNumber(-11))
	fmt.Println("0: ", determineNumber(0))
}
