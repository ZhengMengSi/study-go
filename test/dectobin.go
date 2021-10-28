package main

import (
	"fmt"
	"strconv"
)

func main() {
	decimal := 8
	binary := ""

	for ; decimal > 0; decimal /= 2 {
		lsb := decimal % 2
		binary = strconv.Itoa(lsb) + binary
	}

	//var a float64 = 0.1 + 0.2

	fmt.Println(.1 + .2)
	var a float64 = .1
	var b float64 = .2
	fmt.Println(a + b)
	fmt.Printf("%.54f\n", .1+.2)
}
