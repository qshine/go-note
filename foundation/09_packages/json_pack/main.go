package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var json_string = `{
	"name": "admin",
	"age": 20,
	"contact": {
		"ip": "0.0.0.0",
		"phone": "1234567"
	}
}`

// 解码: 需要将string转换为byte切片, 然后进行反序列化处理
func test_decode_1() {
	type Contact struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Contact struct {
			Ip    string `json:"ip"`
			Phone string `json:"phone"`
		} `json:"contact"`
	}

	// 将json_string字符串反序列化到变量
	var c Contact
	err := json.Unmarshal([]byte(json_string), &c)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Println(c)

	// {admin 20 {0.0.0.0 1234567}}
}

// 解码: 使用map结构
func test_decode_2() {
	var c map[string]interface{}
	err := json.Unmarshal([]byte(json_string), &c)
	if err != nil {
		log.Fatalln("error")
		return
	}
	fmt.Println(c)
	fmt.Println(c["name"])
	fmt.Println(c["contact"])

	// 访问内部的map类型
	// 因为每个键对应值的类型都是interface, 所以需要将值转换为合适的类型才能处理
	fmt.Println(c["contact"].(map[string]interface{})["ip"])

	// map[age:20 contact:map[ip:0.0.0.0 phone:1234567] name:admin]

}

// 编码: 将map类型转换为json
func test_encode_1() {
	c := make(map[string]interface{})
	c["name"] = "admin"
	c["age"] = 20
	c["contact"] = map[string]interface{}{
		"phone":   "12345678",
		"address": "xxx",
	}

	fmt.Println(c)

	// 进行序列化, 返回一个byte切片
	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		log.Fatalln("error")
		return
	}
	fmt.Println(string(data))

	/*
	map[age:20 contact:map[address:xxx phone:12345678] name:admin]
	{
		"age": 20,
		"contact": {
			"address": "xxx",
			"phone": "12345678"
		},
		"name": "admin"
	}
	 */
}

// 编码: 将struct类型转为json
func test_encode_2() {
	// 注意这里必须用大写, 以使json包可以发现
	type person struct {
		Name string
		Age  int
	}

	p := person{
		Name: "admin",
		Age:  20,
	}

	fmt.Println(p)

	data, err := json.Marshal(p)
	if err != nil {
		log.Fatalln("error")
		return
	}
	fmt.Println(string(data))

	/*
	{admin 20}
	{"Name":"admin","Age":20}
	 */
}

func main() {
	//test_decode_1()
	//test_decode_2()
	//test_encode_1()
	test_encode_2()
}
