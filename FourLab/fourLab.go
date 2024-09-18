package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func averageAge(customMap map[string]int) float64 {
	var count, age int
	for _, value := range customMap {
		age += value
		count++
	}
	return float64(age / count)
}

func main() {
	//Задание 1
	fmt.Println("Taske one")
	myMap := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	myMap["Artem"] = 50
	fmt.Println(myMap)

	//Задание 2
	fmt.Println("\nTaks two")
	fmt.Println(averageAge(myMap))

	//Задание 3
	fmt.Println("\nTaks three")
	delete(myMap, "Bob")
	fmt.Println("Map after delete Bob: ", myMap)

	//Задание 4
	fmt.Println("\nTaks four")
	var myString string
	fmt.Print("Enter string for Upper: ")
	fmt.Scanln(&myString)
	myString = strings.ToUpper(myString)
	fmt.Println("After: ", myString)

	//Задание 5
	fmt.Println("\nTaks five")
	fmt.Print("Enter the numbers separated by spaces: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	myString = scanner.Text()
	parts := strings.Fields(myString)

	var sum int
	for _, part := range parts {
		number, err := strconv.Atoi(part)
		if err != nil {
			fmt.Printf("Error with '%s': %v\n", part, err)
			continue
		}
		sum += number
	}

	fmt.Println("Sum: ", sum)

	//Задание 6
	fmt.Println("\nTaks six")
	var n int
	fmt.Print("Enter the array count: ")
	fmt.Scanln(&n)

	fmt.Println("Enter the array element:")
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&arr[i])
	}

	fmt.Println("Reverse massive: ")
	for i := n - 1; i >= 0; i-- {
		fmt.Printf("%d ", arr[i])
	}
}
