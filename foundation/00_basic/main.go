package main

import (
	"fmt"
	"unicode/utf8"
)

func hello() {
	fmt.Println("hello world !")
}

// 变量声明方式
func testVariable() {
	// 声明变量类型, int型初始值是0
	var a, b int
	fmt.Println(a, b)

	// 直接初始化
	var c, d = 123, "hello"
	fmt.Println(c, d)

	// := 这种方式只能在函数内部使用
	e, f := 456, "world"
	fmt.Println(e, f)

	/*
	0 0
	123 hello
	456 world
	 */
}

// 字符串长度
func testString1() {
	var s = "china中国"
	// 得到的是字节长度
	fmt.Println(len(s))

	// 计算字符串的长度
	length := utf8.RuneCountInString(s)
	fmt.Println(length)

	/*
	11
	7
	 */
}

// 字符串 byte 转换
func testString2() {
	var s1 = "hello world"
	var b = []byte(s1) // 字符串转字节切片
	var s2 = string(b) // 字节切片转字符串
	fmt.Println(b)
	fmt.Println(s2)

	/*
	[104 101 108 108 111 32 119 111 114 108 100]
	hello world
	 */
}

// 函数: 带有返回值
func testSwap(x string, y string) (string, string) {
	return y, x
}

func main() {
	//hello()
	//testVariable()
	//testString1()
	//testString2()
	fmt.Println(testSwap("hello", "world"))  // "hello", "world"
}
