package main

import "fmt"

/*
func main() {
	var a *int
	*a = 100
	fmt.Println(*a)
}

/*
该程序会引发panic, 因为Go语言中对于引用类型的变量, 在使用的时候不仅要声明, 还要分配内存空间, 否则就没有办法存储

值类型的声明不需要分配内存空间, 是因为它们在声明的时候已经默认分配好了内存空间
*/

func main() {
	var a *int
	a = new(int)
	*a = 10
	fmt.Println(a)
}
