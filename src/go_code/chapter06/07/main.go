package main

import "fmt"

func main() {
	var arr [5]int = [...]int{5, 9, 7, 2, 1}
	sum := 0
	for _, val := range arr {
		sum += val
	}
	fmt.Printf("和=%d", sum)
}
