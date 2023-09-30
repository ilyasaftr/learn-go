package main

import (
	"fmt"
)

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	// kode di atas adalah sebuah array jika ingin menjadikannya sebuah slice maka hapus ...

	slice := months[4:7]

	// variable slice di atas mengubah array di months menjadi sebuah slice sehingga kita
	// tidak pelru mengganti ... ke kosong tidak apa apa

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	var slice2 = months[2:4]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Ilyasa")
	fmt.Println(slice3)
	slice3[0] = "BukanDesember"
	fmt.Println(slice3)

	fmt.Println("----------")
	fmt.Println("Month 		: ", months)
	fmt.Println("Slice 1 	: ", slice)
	fmt.Println("Slice 2 	: ", slice2)
	fmt.Println("Slice 3 	: ", slice3)
	fmt.Println("----------")

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Ilyasa"
	newSlice[1] = "Fathur Rahman"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	copySlice[0] = "Hello"
	fmt.Println(copySlice)
	fmt.Println(len(copySlice))
	fmt.Println(cap(copySlice))
}
