package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

// 声明并初始化所有字段
func test1() {
	// 这种形式习惯写在一行, 值的顺序非常重要
	book1 := Books{"Go语言", "www.go.com", "Go", 100}
	fmt.Println(book1)

	book2 := Books{
		title:   "Go语言",
		author:  "www.demo.com",
		subject: "Go",
		book_id: 200,
	}
	fmt.Println(book2)

	/*
		{Go语言 www.go.com Go 100}
		{Go语言 www.demo.com Go 200}
	*/
}

// 使用结构类型声明变量
func test2() {
	var Book1 Books
	Book1.title = "Go语言"
	Book1.author = "www.demo.com"
	Book1.subject = "Go"
	Book1.book_id = 300
	fmt.Println(Book1)

	/*
		{Go语言 www.demo.com Go 300}
	*/
}

type user struct {
	name string
	age  int
}

type teacher struct {
	person user // 内嵌结构体
	level  int
}

func test3() {
	a := teacher{
		person: user{
			name: "zhang",
			age:  30,
		},
		level: 10,
	}
	fmt.Println(a)

	/*
		{{zhang 30} 10}
	*/
}

type company struct {
	address string
}

type person struct {
	*company
	name string
	age  int
}

func test4() {
	p := person{
		company: &company{
			address: "xxxx",
		},
		name: "admin",
		age:  20,
	}

	fmt.Println(p)
	fmt.Println(*p.company)
}

func main() {
	// test1()
	// test2()
	// test3()
	test4()
}
