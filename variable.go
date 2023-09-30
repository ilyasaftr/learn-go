package main

import "fmt"

func main() {
	var name string

	name = "Ilyasa"
	fmt.Println(name)

	// Auto Detect Data Type
	var fullName = "Ilyasa Fathur Rahman"
	fmt.Println(fullName)

	var age = 21
	fmt.Println(age)

	// Var tidak wajib jika langsung menginisialisasikan variablenya
	country := "Indonesia"
	fmt.Println(country)

	// Multiple Variable Declaration
	var (
		firstName         = "Ilyasa"
		middleName        = "Fathur"
		lastName   string = "Rahman"
	)

	fmt.Println(firstName, middleName, lastName)
}
