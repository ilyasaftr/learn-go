package main

import "fmt"

func main() {
	type NoKTP string
	type married bool

	var noKtpIlyasa NoKTP = "121004101012812"
	var marriedStatus married = true
	fmt.Println(noKtpIlyasa)
	fmt.Println(marriedStatus)
}
