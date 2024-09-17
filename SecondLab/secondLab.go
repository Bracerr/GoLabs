package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func countAverageNumber(a, b int) float64 {
	return float64((a + b) / 2)
}

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

func stringLenght(a string) int {
	return len(a)
}

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

	//Задание 2
	fmt.Println("\nTask two.")
	fmt.Println("10: ", determineNumber(10))
	fmt.Println("-11: ", determineNumber(-11))
	fmt.Println("0: ", determineNumber(0))

	//Задание 3
	fmt.Println("\nTask three.")
	for i := 0; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	//Задание 4
	fmt.Println("\nTask four.")
	fmt.Println(stringLenght("aaaa"))

	//Задание 5
	fmt.Println("\nTask five.")
	rect := Rectangle{Width: 5, Height: 10}
	fmt.Println("Area: ", rect.Area())

	//Задание 6
	fmt.Println("\nTask six.")
	fmt.Println("Average: ", countAverageNumber(10, 20))
}
