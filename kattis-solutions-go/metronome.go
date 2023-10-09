package main

import "fmt"

func main() {
	var input float32
	fmt.Scanln(&input)
	result := input / 4.0
	fmt.Printf("%.2f", result)
}
