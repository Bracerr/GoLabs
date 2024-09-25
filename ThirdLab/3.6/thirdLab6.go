package main

import "fmt"

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
	fmt.Println("\nTask six")
	strings := []string{"Go", "Gopher", "Программирование", "Самая длинная строка"}
	longest := findLongestString(strings)
	fmt.Println("Longest string: ", longest)
}
