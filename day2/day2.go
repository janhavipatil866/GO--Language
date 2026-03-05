package main

import (
	"fmt"
)

// day2.go: Demonstrate Go basics and common control statements.
func main() {
	fmt.Println("Welcome to Go Day 2!")

	// Variables and types
	var a int = 10
	b := 20.5
	c := "hello"
	fmt.Printf("a=%d b=%.2f c=%s\n", a, b, c)

	// Slice iteration
	nums := []int{1, 2, 3, 4, 5}
	for i, v := range nums {
		fmt.Printf("nums[%d]=%d\n", i, v)
	}

	// Map usage
	info := map[string]string{"name": "Alice", "city": "NY"}
	for k, v := range info {
		fmt.Printf("%s: %s\n", k, v)
	}

	// Control statements examples

	// if-else
	if a > 5 {
		fmt.Println("a is greater than 5")
	} else if a == 5 {
		fmt.Println("a is equal to 5")
	} else {
		fmt.Println("a is less than 5")
	}

	// switch-case
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Another day")
	}

	// for loop with continue and break
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // skip this iteration
		}
		if i == 4 {
			break // exit loop early
		}
		fmt.Println("for i=", i)
	}

	// Calling a simple function
	result := add(a, int(b))
	fmt.Println("add(a,b) =", result)

	// Error handling example
	if err := divide(10, 0); err != nil {
		fmt.Println("divide error:", err)
	}
}

// add returns the sum of two integers
func add(x, y int) int {
	return x + y
}

// divide performs integer division and returns an error on divide by zero
func divide(x, y int) error {
	if y == 0 {
		return fmt.Errorf("cannot divide by zero")
	}
	fmt.Printf("%d / %d = %d\n", x, y, x/y)
	return nil
}
