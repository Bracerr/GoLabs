package main

import "fmt"

func main() {
	myMap := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	delete(myMap, "Bcob")
	fmt.Println("Map after delete Bob: ", myMap)
}
