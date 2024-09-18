package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Person struct {
	name string
	age  int
}

func (p Person) getInfo() {
	fmt.Println("Name: ", p.name, "\nAge: ", p.age)
}

func (p *Person) birthday() {
	p.age++
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	radius float64
}

func (c Circle) getArea() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func countShapeArrayArea(shapes []Shape) {
	for _, shape := range shapes {
		fmt.Printf("Area: %.2f\n", shape.Area())
	}
}

type Stringer interface {
	getBookInfo()
	setPrice(newPrice int)
	setAuthor(newAuthor string)
}

type Book struct {
	id     int
	title  string
	author string
	price  int
}

func (b Book) getBookInfo() {
	fmt.Println("Id: ", b.id)
	fmt.Println("Title: ", b.title)
	fmt.Println("Author: ", b.author)
	fmt.Println("Price: ", b.price)
}

func (b *Book) setPrice(newPrice int) {
	b.price = newPrice
}

func (b *Book) setAuthor(newAuthor string) {
	b.author = newAuthor
}

func main() {
	//Задание 1
	fmt.Println("Task one")
	testPerson := Person{name: "Bob", age: 18}
	testPerson.getInfo()

	//Задание 2
	fmt.Println("\nTask two")
	testPerson.birthday()
	testPerson.getInfo()

	//Задание 3
	fmt.Println("\nTask three")
	testCircle := Circle{radius: 12}
	fmt.Println("Area: ", testCircle.getArea())

	//Задание 4
	fmt.Println("\nTask four")
	rect := Rectangle{Width: 5, Height: 4}
	circle := Circle{radius: 3}
	fmt.Println("Rect area: ", rect.Area())
	fmt.Println("Circle area: ", circle.Area())

	//Задание 5
	fmt.Println("\nTask five")
	shapes := []Shape{rect, circle, testCircle}
	countShapeArrayArea(shapes)

	//Задание 6
	fmt.Println("\nTask six")
	var testInteface Stringer = &Book{
		id:     1,
		title:  "Chipolino",
		author: "",
		price:  122}
	testInteface.getBookInfo()
	testInteface.setAuthor("testAuthor")
	testInteface.setPrice(0)
	fmt.Println()
	testInteface.getBookInfo()
}
