package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var myString string
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
}
