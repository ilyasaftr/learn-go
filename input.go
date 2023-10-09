package main

import "fmt"

func main() {
	fmt.Println("Masukkan nama anda : ")
	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		return
	}
	fmt.Println(name)
}
