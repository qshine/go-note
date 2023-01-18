package main

import (
	"encoding/json"
	"fmt"
)

// struct <-> json
func test0() {
	type T struct {
		Name string
		Age  int
	}

	// struct -> json
	t := T{
		Name: "test",
		Age:  20,
	}

	data, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data)) // {"Name":"test","Age":20}

	// json -> struct
	var t0 T
	err = json.Unmarshal([]byte(data), &t0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t0) // {test 20}
}

// map <-> json
func test1() {
	t := map[string]interface{}{
		"name": "xxx",
		"age":  233333,
	}
	data, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data)) // {"age":233333,"name":"xxx"}

	var res map[string]interface{}
	err = json.Unmarshal([]byte(data), &res)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res) // map[age:233333 name:xxx]
}

func main() {
	//test0()
	test1()
}
