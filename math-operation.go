package main

import "fmt"

func main() {
	// Basic Math Operation
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c)

	// Augmented Assignments
	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	// Unary Operator
	i++ // i = i + 1
	fmt.Println(i)

	var negative = -100
	var positive = -100

	fmt.Println(negative)
	fmt.Println(positive)
}
