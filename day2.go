package main

import (
	"fmt"
)

// day2.go demonstrates basic Go concepts such as functions,
// slices, maps, loops, and simple error handling.
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

	// Calling a function
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
