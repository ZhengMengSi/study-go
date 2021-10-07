package main

import "fmt"

func main() {
	// 第一种使用方式
	/*var a map[string]string
	a = make(map[string]string, 10)
	a["no1"] = "松江"
	a["no2"] = "无用"
	a["no1"] = "武松"
	a["no3"] = "666"
	fmt.Println(a)*/

	// 第二种使用方式
	/*cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)*/

	// 第三种使用方式
	/*heros := map[string]string{
		"hero1": "宋江",
		"hero2": "卢俊义",
		"hero3": "吴用",
	}
	fmt.Println(heros)*/

	// value是map
	studentMap := make(map[string]map[string]string)

	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "北京"

	studentMap["stu02"] = make(map[string]string, 3)
	studentMap["stu02"]["name"] = "mary"
	studentMap["stu02"]["sex"] = "男"
	studentMap["stu02"]["address"] = "北京"

	/*for k1, v1 := range studentMap {
		fmt.Println("k1=", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%v v2=%v \n", k2, v2)
		}
		fmt.Println()
	}*/

	fmt.Println(len(studentMap))
}
