package main

import "fmt"

func main() {
	// 声明类型
	var a map[string]int
	// 初始化
	a = make(map[string]int)
	a["age"] = 20

	fmt.Println(a)

	b := make(map[string]string)
	b["name"] = "xxx"
	fmt.Println(b)

	c := map[string]string{
		"address": "asia",
	}
	fmt.Println(c)

}

/*
map[age:20]
map[name:xxx]
map[address:asia]
*/
