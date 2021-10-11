package main

import "fmt"

func main() {
	// map的声明
	var a map[string]string
	// 在使用map前，需要先make，make的作用就是给map分配数据空间
	a = make(map[string]string, 10)
	a["no2"] = "吴用"
	a["no1"] = "宋江"
	a["no1"] = "李逵"
	a["no3"] = "吴用"
	fmt.Println(a)
	for k, v := range a {
		fmt.Printf("k: %v,v: %v \n", k, v)
	}
}
