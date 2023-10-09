package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string
	person := map[string]string{
		"name":    "Ilyasa",
		"address": "Bogor",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Ilyasa"
	book["ups"] = "Salah"

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))

}
