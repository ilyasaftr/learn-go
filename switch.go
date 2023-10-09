package main

import "fmt"

func main() {
	name := "Ilyasa"

	switch name {
	case "Ilyasa":
		fmt.Println("Hello Ilyasa")
	case "Fathur":
		fmt.Println("Hello Fathur")
	default:
		fmt.Println("Boleh kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama terlalu pendek")
	}

	length2 := len(name)
	switch {
	case length2 > 30:
		fmt.Println("Nama terlalu panjang 2")
	case length2 > 5:
		fmt.Println("Nama lumayan panjang 2")
	default:
		fmt.Println("Nama sudah benar 2")
	}
}
