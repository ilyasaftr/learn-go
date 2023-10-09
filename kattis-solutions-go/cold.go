package main

import "fmt"

func main() {
	var dataCount, dataCountTemp, i uint8
	fmt.Scanln(&dataCount)
	for i = 0; i < dataCount; i++ {
		var temp int32
		fmt.Scan(&temp)
		if temp < 0 {
			dataCountTemp++
		}
	}
	fmt.Println(dataCountTemp)
}
