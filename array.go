package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Ilyasa"
	names[1] = "Fathur"
	names[2] = "Rahman"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{1, 2, 3}
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	var values2 = [...]int{100, 200, 300}
	fmt.Println(values2[0])
	fmt.Println(values2[1])
	fmt.Println(values2[2])

	fmt.Println(len(values2))
	fmt.Println(values2[0])
	values2[0] = 2000
	fmt.Println(values2[0])
}
