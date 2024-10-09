package main

import "fmt"

func main() {
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
