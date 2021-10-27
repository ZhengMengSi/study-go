package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "1,2,3"
	//b := a[1:3]
	c := strings.Split(a, ",")
	d := strings.SplitN(a, ",", 2)
	fmt.Println(c)
	fmt.Println(d[0])
	fmt.Println(d[1])
}
