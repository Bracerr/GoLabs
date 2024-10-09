package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("\nTaks four")
	var myString string
	fmt.Print("Enter string for Upper: ")
	fmt.Scanln(&myString)
	myString = strings.ToUpper(myString)
	fmt.Println("After: ", myString)
}
