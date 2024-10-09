package main

import "fmt"

func averageAge(customMap map[string]int) float64 {
	var count, age int
	for _, value := range customMap {
		age += value
		count++
	}
	return float64(age / count)
}

func main() {
	myMap := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	fmt.Println("\nTaks two")
	fmt.Println(averageAge(myMap))
}
