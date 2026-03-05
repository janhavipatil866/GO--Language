package main

import "fmt"

// 1️⃣ Function Example
func add(a int, b int) int {
	return a + b
}

// 2️⃣ Factorial Function
func factorial(n int) int {
	fact := 1
	for i := 1; i <= n; i++ {
		fact = fact * i
	}
	return fact
}

// 5️⃣ Even/Odd Function
func checkEvenOdd(num int) {
	if num%2 == 0 {
		fmt.Println(num, "is Even")
	} else {
		fmt.Println(num, "is Odd")
	}
}

func main() {

	fmt.Println("----- Function Example -----")
	result := add(10, 20)
	fmt.Println("Sum:", result)

	fmt.Println("\n----- Factorial Example -----")
	fmt.Println("Factorial of 5:", factorial(5))

	fmt.Println("\n----- Array Example -----")

	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	fmt.Println("\n----- Slice Example -----")

	sliceNumbers := []int{10, 20, 30, 40}
	sliceNumbers = append(sliceNumbers, 50)

	for _, value := range sliceNumbers {
		fmt.Println(value)
	}

	fmt.Println("\n----- Even Odd Example -----")
	checkEvenOdd(10)

}
