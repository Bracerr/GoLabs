package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	user := user{name: "Ivan", age: 19}
	userPointer := &user
	(*userPointer).age += 1
	fmt.Println(user.age)
}
