package main

import (
	"fmt"
	"go_code/chapter06/01fundemo/utils"
)

func main() {
	var n1 float64 = 1.2
	var n2 float64 = 2.3
	var operator byte = '-'
	var res float64 = utils.Cal(n1, n2, operator)
	fmt.Println("res=", res)
}
