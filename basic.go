package main

import "fmt"

func ar_crc(no1, no2 int) float64 {
	const pi = 3.14
	return pi * float64(no1*no2)
}
func main() {
	i := 0
	for i < 2 {
		fmt.Println("Hello, Duniya")
		i++
	}
	for i = 0; i < 3; i++ {
		fmt.Println("Hello, Duniya")
	}

	fmt.Print("\nCircle Area(4,5)\n")
	var h, r int
	h = 4
	r = 5
	res := ar_crc(h, r) //same as [var res int =func()]
	fmt.Print("\nSol:- ", res)
}
