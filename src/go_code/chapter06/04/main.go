package main

import "fmt"

func main() {
	res := fbn(6)
	fmt.Println(res)
}

func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fbn(n-1) + fbn(n-2)
}
