package main

import "fmt"

func main() {
	res1, res2 := getSumAndSub(1, 2)
	fmt.Printf("res1=%v,res2=%v", res1, res2)
}

func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}
