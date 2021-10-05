package main

import "fmt"

func main() {
	res := fbn(1)
	fmt.Println(res)
}

func fbn(n int) int {
	if n == 10 {
		return 1
	}
	return (fbn(n+1) + 1) * 2
}
