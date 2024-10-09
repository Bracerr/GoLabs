package main

import "fmt"

func main() {
	fmt.Println("Taske one")
	myMap := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	myMap["Artem"] = 50
	fmt.Println(myMap)
}
