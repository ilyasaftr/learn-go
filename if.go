package main

import "fmt"

func main() {
	var name = "Fathur"

	if name == "Ilyasa" {
		fmt.Println("Hallo Ilyasa")
	} else if name == "Fathur" {
		fmt.Println("Hallo Fathur")
	} else {
		fmt.Println("Hi, kenalan donk")
	}

	// if short statement
	// var length = len(name); become ->
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
