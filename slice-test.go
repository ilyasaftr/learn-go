package main

import "fmt"

func main() {
	// slice months
	var months = []string{
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

	// perbedaan mengakses menggunakan array dan slices
	var arrayData = months[0]
	var sliceData = months[0:1]

	// kesimpulan adalah jika menggunakan arrayData atau mengakses dengan array saja
	// maka itu hanya mengcopy value tersebut tidak melakukan referensi sehingga jika diubah maka
	// data asli tidak keubah

	// sedangkan jika mengakses menggunakan sliceData atau mengakses dengan slice itu akan
	// melakukan referensi memori sehingga ketika slice diubah maka akan ikut sampai ke data asli

	fmt.Println("-----------")
	fmt.Println(arrayData)
	fmt.Println(months)
	arrayData = "Ini Januari"
	fmt.Println(arrayData)
	fmt.Println(months)

	fmt.Println("-----------")
	fmt.Println(sliceData)
	sliceData[0] = "Ini Januari"
	fmt.Println(sliceData)
	fmt.Println(months)
}
