package main

import "fmt"

func main() {

	// -----------------------------
	// 1️⃣ Variables & Data Types
	// -----------------------------

	// Integer types
	var intValue int = 10
	var int8Value int8 = 8
	var int16Value int16 = 160
	var int32Value int32 = 320
	var int64Value int64 = 640

	// Unsigned integer
	var uintValue uint = 20

	// Float types
	var float32Value float32 = 10.5
	var float64Value float64 = 99.99

	// Boolean
	var isLearning bool = true

	// String
	var name string = "Janhavi"

	fmt.Println("----- Data Types in Go -----")
	fmt.Println("int:", intValue)
	fmt.Println("int8:", int8Value)
	fmt.Println("int16:", int16Value)
	fmt.Println("int32:", int32Value)
	fmt.Println("int64:", int64Value)
	fmt.Println("uint:", uintValue)
	fmt.Println("float32:", float32Value)
	fmt.Println("float64:", float64Value)
	fmt.Println("bool:", isLearning)
	fmt.Println("string:", name)

	// -----------------------------
	// 2️⃣ Short Variable Declaration
	// -----------------------------

	age := 22
	city := "Pune"

	fmt.Println("\nShort Declaration:")
	fmt.Println("Age:", age)
	fmt.Println("City:", city)

	// -----------------------------
	// 3️⃣ Constants
	// -----------------------------

	const pi float64 = 3.14
	const country = "India"

	fmt.Println("\nConstants:")
	fmt.Println("PI:", pi)
	fmt.Println("Country:", country)

	// -----------------------------
	// 4️⃣ Type Conversion
	// -----------------------------

	var num1 int = 10
	var num2 float64 = float64(num1) // int to float64

	fmt.Println("\nType Conversion:")
	fmt.Println("Converted int to float64:", num2)

	var num3 float64 = 9.8
	var num4 int = int(num3) // float64 to int

	fmt.Println("Converted float64 to int:", num4)
}