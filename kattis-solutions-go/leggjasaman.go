package main

import "fmt"

func main() {
	var number1, number2 uint16
	fmt.Scanln(&number1)
	fmt.Scanln(&number2)
	fmt.Println(number1 + number2)
}
