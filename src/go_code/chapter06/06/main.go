package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var numArr = [...]int{1: 800, 0: 900}

func init() {
	var intArr3 [5]int
	len := len(intArr3)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		intArr3[i] = rand.Intn(100)
	}
	fmt.Println("交换前=", intArr3)

	temp := 0
	for i := 0; i < len/2; i++ {
		temp = intArr3[len-1-i]
		intArr3[len-1-i] = intArr3[i]
		intArr3[i] = temp
	}

	fmt.Println("交换后=", intArr3)
}

func main() {
	funcvar := makeSuffix(".jpg")
	fmt.Println(funcvar("1"))
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
