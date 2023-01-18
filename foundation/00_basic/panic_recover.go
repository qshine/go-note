package main

/*
异常处理
 */

import "fmt"

func test_1() {
	// recover函数 只能在defer中使用
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("444")
		}
	}()

	fmt.Println("111")
	panic("err demo")
	fmt.Println("222") // 此行不输出

	/*
	111
	444
	 */
}

func Call() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("333")
		}
	}()
	panic("error test")
}

func test_2() {
	fmt.Println("111")
	Call()
	fmt.Println("222")

	/*
	111
	333
	222
	 */
}

func main() {
	//test_1()
	test_2()
}
