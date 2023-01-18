package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 写入文件. Write 和 WriteString
func test0() {
	file, err := os.OpenFile("test0.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	str := "hello world!"
	file.Write([]byte(str))     // 写入字节切片数据
	file.WriteString("world !") // 直接写入字符串
}

// 先写入缓存
func test1() {
	file, err := os.OpenFile("test1.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()

	// 建立缓存区
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello沙河\n") //将数据先写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件
}

// ioutil
func test2() {
	str := "hello world"
	err := ioutil.WriteFile("test2.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

func main() {
	test0()
	test1()
	test2()
}
