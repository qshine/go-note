package main

import "fmt"

/*
Map是一种无序的键值对集合, 使用以下方式来定义Map

1. map_variable := make(map[key_type]value_type, capacity)
capacity表示可以存储指定个数键值对

2. map_variable := make(map[key_type]value_type)

3. map[key_type]value_type{key1:value1, key2:value2}

注意:
1. map不是并发安全的
2. 在函数之间传递 map 不是传递的副本, 而是传递的引用, 和slice类似
*/

func test1() {
	// 使用make创建
	country := make(map[string]string)
	country["china"] = "北京"
	country["india"] = "新德里"

	for k, v := range country {
		fmt.Println(k, v)
	}

	// 检测key在不在
	capital, ok := country["美国"]
	if ok {
		fmt.Println("美国首都: ", capital)
	} else {
		fmt.Println("没有美国")
	}
}

func test_delete() {
	// 直接初始化
	map_string := map[string]string{"a": "b", "c": "d"}
	fmt.Println(map_string)
	// 删除某个元素
	delete(map_string, "a")
	fmt.Println(map_string)
}

func cnt_words() {
	// 统计词频
	words := []string{"a", "b", "a", "c", "d", "c"}
	results := make(map[string]int)
	for _, word := range words {
		if nums, ok := results[word]; ok {
			results[word] = nums + 1
		} else {
			results[word] = 1
		}
	}
	fmt.Println(results)
}

func main() {
	test1()
	fmt.Println()
	test_delete()
	fmt.Println()
	cnt_words()
}

/*
china 北京
india 新德里
没有美国

map[a:b c:d]
map[c:d]

map[a:2 b:1 c:2 d:1]
*/
