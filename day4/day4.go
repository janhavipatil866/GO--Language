package main

import (
	"errors"
	"fmt"
)

// 1️⃣ Struct Example
type Student struct {
	Name string
	Age  int
	City string
}

func studentExample() {
	s := Student{
		Name: "Janhavi",
		Age:  22,
		City: "Pune",
	}

	fmt.Println("Student Details:")
	fmt.Println("Name:", s.Name)
	fmt.Println("Age:", s.Age)
	fmt.Println("City:", s.City)
}

// 2️⃣ Struct with Method
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func rectangleExample() {
	r := Rectangle{Width: 10, Height: 5}
	fmt.Println("Rectangle Area:", r.Area())
}

// 3️⃣ Interface Example
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func interfaceExample() {
	var s Shape
	s = Circle{Radius: 7}

	fmt.Println("Circle Area:", s.Area())
}

// 4️⃣ Map Example
func mapExample() {

	languages := map[int]string{
		1: "Go",
		2: "Java",
		3: "Python",
		4: "JavaScript",
	}

	fmt.Println("Programming Languages:")
	for key, value := range languages {
		fmt.Println(key, "->", value)
	}
}

// 5️⃣ Error Handling Example
func divide(a int, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

func errorExample() {

	result, err := divide(10, 0)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}

// Main Function
func main() {

	fmt.Println("----- Struct Example -----")
	studentExample()

	fmt.Println("\n----- Method Example -----")
	rectangleExample()

	fmt.Println("\n----- Interface Example -----")
	interfaceExample()

	fmt.Println("\n----- Map Example -----")
	mapExample()

	fmt.Println("\n----- Error Handling Example -----")
	errorExample()
}
